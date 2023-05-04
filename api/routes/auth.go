package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/cookieAuthentication/api/handlers"
)

func attachAuthRoutes(app *fiber.App) {
	router := app.Group("auth")
	router.Post("/register", handlers.Register)
	router.Post("/login", handlers.Login)
	router.Get("/user", handlers.GetUser)
}
