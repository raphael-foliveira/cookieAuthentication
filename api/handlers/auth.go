package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/cookieAuthentication/api/auth"
	"github.com/raphael-foliveira/cookieAuthentication/api/database"
)

type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var userCreate UserCreate
	err := c.BodyParser(&userCreate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Data",
		})
	}
	newUser, err := database.CreateUser(userCreate.Username, userCreate.Password, userCreate.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func Login(c *fiber.Ctx) error {
	var userLoginData UserLogin
	fmt.Println("handling user login")

	err := c.BodyParser(&userLoginData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user credentials",
		})
	}
	user, err := database.FindUserByUsername(userLoginData.Username)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid username",
		})
	}
	if user.Password != userLoginData.Password {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}
	sessionToken, err := auth.RenewSessionToken(user.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not create session token",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    sessionToken.Token,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
	})
	return c.Status(200).JSON(fiber.Map{
		"message": "Login successful",
	})
}

func GetUser(c *fiber.Ctx) error {
	userCookie := c.Cookies("auth")
	if userCookie == "" {
		fmt.Println("no cookie")
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "User-Logged-in",
		Value:    "true",
		SameSite: "Lax",
	})
	user, err := database.FindUserBySessionToken(userCookie)
	if err != nil {
		fmt.Println(err)
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return c.Status(200).JSON(user)
}
