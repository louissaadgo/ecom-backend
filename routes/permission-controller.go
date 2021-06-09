package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/models"
)

func permissions(c *fiber.Ctx) error {
	var permissions []models.Permission

	query, err := database.DB.Query(`SELECT * FROM permissions;`)
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	defer query.Close()
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	for query.Next() {
		var permission models.Permission
		err := query.Scan(&permission.ID, &permission.Name)
		if err != nil {
			return fiber.NewError(400, "Something went wrong")
		}
		permissions = append(permissions, permission)
	}
	return c.JSON(permissions)
}
