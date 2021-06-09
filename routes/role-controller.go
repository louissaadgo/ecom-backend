package routes

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/models"
)

func roles(c *fiber.Ctx) error {
	var roles []models.Role

	query, err := database.DB.Query(`SELECT * FROM roles;`)
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	defer query.Close()
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	for query.Next() {
		var role models.Role
		err := query.Scan(&role.ID, &role.Name)
		if err != nil {
			return fiber.NewError(400, "Something went wrong")
		}
		roles = append(roles, role)
	}
	return c.JSON(roles)
}

func createRole(c *fiber.Ctx) error {
	var role models.Role
	if err := c.BodyParser(&role); err != nil {
		return fiber.NewError(400, "Failed to create a user")
	}
	_, err := database.DB.Exec(`INSERT INTO roles(name,permissions) VALUES($1,$2)`, role.Name, pq.Array(role.Permissions))
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(400, "Something went wrong")
	}
	return c.SendString("Success")
}

func getRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	var role models.Role
	q := database.DB.QueryRow(`SELECT id, name FROM roles WHERE id = $1`, id)
	err = q.Scan(&role.ID, &role.Name)
	if err == sql.ErrNoRows {
		return fiber.NewError(400, "Something went wrong")
	}
	return c.JSON(role)
}

func updateRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return fiber.NewError(400, "something went wrong")
	}
	role.ID = uint(id)

	database.DB.Exec(`UPDATE roles SET name = $1, permissions = $2 WHERE id = $3;`, role.Name, pq.Array(role.Permissions), role.ID)

	return c.JSON(role)
}

func deleteRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	roleID := uint(id)
	database.DB.Exec(`DELETE FROM roles WHERE id = $1;`, roleID)
	return c.SendString("Success")
}
