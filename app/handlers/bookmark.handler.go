package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type bookmarkHandler struct {
	DB       *sqlx.DB
	Validate *validator.Validate
}

func NewBookmarkHandler(db *sqlx.DB, validate *validator.Validate) *bookmarkHandler {
	return &bookmarkHandler{
		DB:       db,
		Validate: validate,
	}
}

func (bh *bookmarkHandler) GetBookmarks(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "get bookmarks"})
}
