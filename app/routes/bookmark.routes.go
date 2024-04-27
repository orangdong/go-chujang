package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/orangdong/go-chujang/app/handlers"
)

func BookmarkRoutes(app *fiber.App, db *sqlx.DB) {
	bookmarks := app.Group("/bookmarks")
	validate := validator.New(validator.WithRequiredStructEnabled())
	bookmarkHandler := handlers.NewBookmarkHandler(db, validate)

	bookmarks.Get("/", bookmarkHandler.GetBookmarks)
	bookmarks.Get("/:id", bookmarkHandler.GetBookmark)
	bookmarks.Post("/", bookmarkHandler.CreateBookmark)
	bookmarks.Put("/:id", bookmarkHandler.UpdateBookmark)
}
