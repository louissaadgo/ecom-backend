package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/models"
)

func createUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(400, "Failed to create a user")
	}
	user.SetPassword(user.Password)

	_, err := database.DB.Exec(`INSERT INTO users(firstName, lastName, email, password) VALUES($1, $2, $3, $4)`, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	return c.SendString("Success")
}
