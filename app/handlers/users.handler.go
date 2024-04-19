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

type userHandler struct {
	DB       *sqlx.DB
	Validate *validator.Validate
}

func NewUserHandler(db *sqlx.DB, validate *validator.Validate) *userHandler {
	return &userHandler{DB: db, Validate: validate}
}

func (u *userHandler) GetUsers(c *fiber.Ctx) error {
	// get users
	users := []entities.UserModel{}
	getUsers := sq.Select("*").From("users")
	query, args, _ := getUsers.ToSql()
	u.DB.Select(&users, query, args...)

	return c.Status(200).JSON(utils.NewSuccessResponse("users fetched successfully", users))
}

func (u *userHandler) GetUserById(c *fiber.Ctx) error {
	// get user by id
	user := entities.UserModel{}
	id := c.Params("id")
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	getUser := psql.Select("*").From("users").Where(sq.Eq{"id": id})
	query, args, e := getUser.ToSql()

	if e != nil {
		return c.Status(400).JSON(utils.NewErrorResponse("invalid id"))
	}

	u.DB.Get(&user, query, args...)

	if user.ID == "" {
		return c.Status(404).JSON(utils.NewErrorResponse("user with id " + id + " not found"))
	}

	return c.Status(200).JSON(utils.NewSuccessResponse("user with id "+id+" fetched successfully", user))
}

func (u *userHandler) CreateUser(c *fiber.Ctx) error {
	// create user
	user := entities.UserCreate{}
	uuid := uuid.New()
	err := c.BodyParser(&user)
	errValidate := utils.Validate(user, u.Validate)

	if err != nil {
		return c.Status(400).JSON(utils.NewErrorResponse("invalid request body"))
	}

	if errValidate != "" {
		return c.Status(400).JSON(utils.NewErrorResponse(errValidate))
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	createUser := psql.Insert("users").Columns("id", "name", "username", "email", "avatar").Values(uuid.String(), user.Name, user.Username, user.Email, user.Avatar)
	query, args, _ := createUser.ToSql()
	_, err = u.DB.Exec(query, args...)

	if err != nil {
		return c.Status(500).JSON(utils.NewErrorResponse("failed to execute SQL query"))
	}

	return c.Status(201).JSON(utils.NewSuccessResponse("user created successfully", user))
}
