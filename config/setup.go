package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/orangdong/go-chujang/app/routes"
)

func SetupAndRunApp() error {
	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(favicon.New())
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	// setup routes
	routes.SetupRoutes(app)

	// attach swagger
	// config.AddSwaggerRoutes(app)

	// get the port and start
	port := "3000"
	app.Listen(":" + port)
	return nil
}
