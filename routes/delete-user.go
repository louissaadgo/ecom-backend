package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
)

func deleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	userID := uint(id)
	database.DB.Exec(`DELETE FROM users WHERE id = $1;`, userID)
	return c.SendString("Success")
}
