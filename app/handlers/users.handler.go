package handlers

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/orangdong/go-chujang/app/entities"
	"github.com/orangdong/go-chujang/app/utils"
)

type userHandler struct {
	DB *sqlx.DB
}

func NewUserHandler(db *sqlx.DB) *userHandler {
	return &userHandler{DB: db}
}

func (u *userHandler) GetUsers(c *fiber.Ctx) error {
	// get users
	q := []entities.UserModel{}
	users := []entities.UserDTO{}
	getUsers := sq.Select("*").From("users")
	query, args, _ := getUsers.ToSql()
	u.DB.Select(&q, query, args...)

	for _, user := range q {
		users = append(users, entities.UserDTO{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Username: user.Username,
			Avatar:   user.Avatar,
		})
	}

	return c.Status(200).JSON(utils.NewSuccessResponse("users fetched successfully", users))
}

func (u *userHandler) GetUserById(c *fiber.Ctx) error {
	// get user by id
	q := entities.UserModel{}
	id := c.Params("id")
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	getUser := psql.Select("*").From("users").Where(sq.Eq{"id": id})
	query, args, e := getUser.ToSql()

	if e != nil {
		return c.Status(400).JSON(utils.NewErrorResponse("invalid id"))
	}

	u.DB.Get(&q, query, args...)

	if q.ID == 0 {
		return c.Status(404).JSON(utils.NewErrorResponse("user with id " + id + " not found"))
	}

	user := &entities.UserDTO{
		ID:       q.ID,
		Name:     q.Name,
		Email:    q.Email,
		Username: q.Username,
		Avatar:   q.Avatar,
	}

	return c.Status(200).JSON(utils.NewSuccessResponse("user with id "+id+" fetched successfully", user))
}

func (u *userHandler) CreateUser(c *fiber.Ctx) error {
	// create user
	user := entities.UserCreate{}
	err := c.BodyParser(&user)
	fmt.Println(user, err)
	// createUser := sq.Insert("users").Columns("name", "email").Values(user.Name, user.Email)
	// query, _, _ := createUser.ToSql()
	// u.DB.MustExec(query)

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "user created successfully", "data": user})
}
