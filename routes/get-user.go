package routes

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/models"
)

func getUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	var user models.User
	q := database.DB.QueryRow(`SELECT id, firstName, lastName, email, role_id FROM users WHERE id = $1`, id)
	err = q.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.RoleID)
	if err == sql.ErrNoRows {
		return fiber.NewError(400, "Something went wrong")
	}
	return c.JSON(user)
}
