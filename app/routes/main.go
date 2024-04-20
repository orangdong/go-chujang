package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "success", "message": "go-chujang ready!"})
	})

	UserRoutes(app, db)
	BookmarkRoutes(app, db)
}
