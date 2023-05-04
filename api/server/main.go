package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raphael-foliveira/cookieAuthentication/api/database"
	"github.com/raphael-foliveira/cookieAuthentication/api/routes"
)

func Start() {
	database.Connect()
	app := fiber.New()
	routes.GetAll(app)
	app.Listen(":8000")
}
