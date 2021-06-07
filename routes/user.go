package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/middlewares"
)

func user(c *fiber.Ctx) error {

	cookie := c.Cookies("JWT")

	if cookie == "" {
		return fiber.NewError(400, "User Not Authenticated")
	}

	id, err := middlewares.ParseJwt(cookie)

	if err != nil {
		return fiber.NewError(400, "Somwthing went wrong")
	}

	query := database.DB.QueryRow(`SELECT firstName FROM users WHERE id = $1;`, id)
	var username string
	if err := query.Scan(&username); err == sql.ErrNoRows {
		return fiber.NewError(400, "Somwthing went wrong")
	}

	return c.JSON(username)
}
