package handlers

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/orangdong/go-chujang/app/entities"
	"github.com/orangdong/go-chujang/app/utils"
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
	// get bookmarks
	bookmarks := []entities.BookmarkModel{}
	getBookmarks := sq.Select("*").From("bookmarks")
	query, args, _ := getBookmarks.ToSql()
	bh.DB.Select(&bookmarks, query, args...)

	return c.Status(200).JSON(utils.NewSuccessResponse("bookmarks fetched successfully", bookmarks))
}

func (bh *bookmarkHandler) GetBookmark(c *fiber.Ctx) error {
	// get bookmark
	bookmark := entities.BookmarkModel{}
	id := c.Params("id")
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	getBookmark := psql.Select("*").From("bookmarks").Where(sq.Eq{"id": id})
	query, args, e := getBookmark.ToSql()
	if e != nil {
		return c.Status(400).JSON(utils.NewErrorResponse("invalid id"))
	}
	bh.DB.Get(&bookmark, query, args...)

	if bookmark.ID == "" {
		return c.Status(404).JSON(utils.NewErrorResponse("bookmark with id " + id + " not found"))
	}

	return c.Status(200).JSON(utils.NewSuccessResponse("bookmark fetched successfully", bookmark))
}

func (bh *bookmarkHandler) CreateBookmark(c *fiber.Ctx) error {
	// create bookmark
	bookmark := entities.BookmarkCreate{}
	bookmarkId := uuid.New()
	userId := "550e8400-e29b-41d4-a716-446655440001"

	err := c.BodyParser(&bookmark)
	errValidate := utils.Validate(bookmark, bh.Validate)
	if err != nil {
		return c.Status(400).JSON(utils.NewErrorResponse("invalid request body"))
	}

	if errValidate != "" {
		return c.Status(400).JSON(utils.NewErrorResponse(errValidate))
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	insertBookmark := psql.Insert("bookmarks").Columns("id", "user_id", "title", "summary", "url", "keywords", "is_private").Values(bookmarkId.String(), userId, bookmark.Title, bookmark.Summary, bookmark.URL, bookmark.KeyWords, bookmark.IsPrivate)
	query, args, _ := insertBookmark.ToSql()
	_, err = bh.DB.Exec(query, args...)

	if err != nil {
		return c.Status(500).JSON(utils.NewErrorResponse("failed to execute SQL query"))
	}

	return c.Status(201).JSON(utils.NewSuccessResponse("bookmark created successfully", map[string]string{"id": bookmarkId.String(), "url": bookmark.URL}))
}

func (bh *bookmarkHandler) UpdateBookmark(c *fiber.Ctx) error {
	// update bookmark
	bookmark := entities.BookmarkUpdate{}
	updatedBookmark := entities.BookmarkModel{}
	id := c.Params("id")
	err := c.BodyParser(&bookmark)
	errValidate := utils.Validate(bookmark, bh.Validate)
	if err != nil {
		return c.Status(400).JSON(utils.NewErrorResponse("invalid request body"))
	}

	if errValidate != "" {
		return c.Status(400).JSON(utils.NewErrorResponse(errValidate))
	}

	updateFields := utils.UpdatedFieldsMap(bookmark)

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	updateBookmark := psql.Update("bookmarks").SetMap(updateFields).Where(sq.Eq{"id": id})
	query, args, _ := updateBookmark.ToSql()
	_, err = bh.DB.Exec(query, args...)

	if err != nil {
		return c.Status(500).JSON(utils.NewErrorResponse("failed to execute SQL query"))
	}

	getUpdatedBookmark := psql.Select("*").From("bookmarks").Where(sq.Eq{"id": id})
	query, args, _ = getUpdatedBookmark.ToSql()
	err = bh.DB.Get(&updatedBookmark, query, args...)
	if err != nil {
		return c.Status(500).JSON(utils.NewErrorResponse("failed to execute SQL query"))
	}

	return c.Status(200).JSON(utils.NewSuccessResponse("bookmark updated successfully", updatedBookmark))
}
