package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/models"
)

func updateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	var user models.User

	user.Id = uint(id)

	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(400, "something went wrong")
	}

	database.DB.Exec(`UPDATE users SET firstName = $1, lastName = $2, email = $3, role_id = $4 WHERE id = $5;`, user.FirstName, user.LastName, user.Email, user.RoleID, user.Id)

	return c.JSON(user)
}
