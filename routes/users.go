package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/models"
)

func users(c *fiber.Ctx) error {
	var users []models.User

	query, err := database.DB.Query(`SELECT * FROM users;`)
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	defer query.Close()
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	for query.Next() {
		var user models.User
		err := query.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.RoleID)
		if err != nil {
			return fiber.NewError(400, "Something went wrong")
		}
		users = append(users, user)
	}
	return c.JSON(users)
}
