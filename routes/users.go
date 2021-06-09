package routes

import (
	"database/sql"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/models"
)

func users(c *fiber.Ctx) error {
	var users []models.User
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	limit := 10
	offset := (page - 1) * limit
	var total int64

	query, err := database.DB.Query(`SELECT * FROM users LIMIT $1 OFFSET $2;`, limit, offset)
	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}
	defer query.Close()
	for query.Next() {
		var user models.User
		err := query.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.RoleID)
		if err != nil {
			return fiber.NewError(400, "Something went wrong")
		}
		users = append(users, user)
	}
	q := database.DB.QueryRow(`SELECT COUNT(*) FROM users;`)
	err = q.Scan(&total)
	if err == sql.ErrNoRows {
		return fiber.NewError(400, "Something went wrong")
	}
	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)),
		},
	})
}
