package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/cookieAuthentication/api/middleware"
)

func GetAll(app *fiber.App) {
	middleware.Attach(app)
	attachAuthRoutes(app)
	attachPostRoutes(app)
}
