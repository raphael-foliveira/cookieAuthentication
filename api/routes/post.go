package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/cookieAuthentication/api/middleware"
)

func attachPostRoutes(app *fiber.App) {
	router := app.Group("posts", middleware.Authorization)
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("post router")
	})
}
