package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/orangdong/go-chujang/app/handlers"
)

func UserRoutes(app *fiber.App, db *sqlx.DB) {

	users := app.Group("/users")
	userHandler := handlers.NewUserHandler(db)
	users.Get("/", userHandler.GetUsers)
	users.Post("/", userHandler.CreateUser)
	users.Get("/:id", userHandler.GetUserById)
}
