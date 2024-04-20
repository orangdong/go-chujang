package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/orangdong/go-chujang/app/handlers"
)

func UserRoutes(app *fiber.App, db *sqlx.DB) {

	users := app.Group("/users")
	validate := validator.New(validator.WithRequiredStructEnabled())
	userHandler := handlers.NewUserHandler(db, validate)

	users.Get("/", userHandler.GetUsers)
	users.Post("/", userHandler.CreateUser)
	users.Get("/:id", userHandler.GetUserById)
	users.Put("/:id", userHandler.UpdateUser)
}
