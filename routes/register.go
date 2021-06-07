package routes

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/models"
)

func register(c *fiber.Ctx) error {

	user := models.User{}
	var data map[string]string

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if user.Password != data["PasswordConfirm"] {
		return fiber.NewError(400, "Invalid password")
	}

	hashedPassword := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(hashedPassword[:])

	_, err := database.DB.Query(`INSERT INTO users(firstName, lastName, email, password) VALUES($1,$2,$3,$4);`, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		errString := fmt.Sprintln(err)
		return fiber.NewError(400, errString)
	}

	return c.JSON(user)
}
