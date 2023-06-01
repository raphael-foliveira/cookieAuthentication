package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Authorization(c *fiber.Ctx) error {
	authCookie := c.Cookies("auth")
	fmt.Println(authCookie)
	c.Next()
	return nil
}
