package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/raphael-foliveira/cookieAuthentication/api/database"
	"github.com/raphael-foliveira/cookieAuthentication/api/handlers"
)

func Start() {
	database.Start()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Get("/user", handlers.GetUser)

	app.Listen(":8000")
}
