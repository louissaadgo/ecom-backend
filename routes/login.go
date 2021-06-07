package routes

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/models"
)

func login(c *fiber.Ctx) error {

	var user models.User
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	query := database.DB.QueryRow(`SELECT id, password FROM users WHERE email = $1;`, data["Email"])
	q := query.Scan(&user.Id, &user.Password)

	if q == sql.ErrNoRows {
		return fiber.NewError(400, "No user found")
	}

	hashedPassword := sha256.Sum256([]byte(data["Password"]))
	newPassword := hex.EncodeToString(hashedPassword[:])

	if user.Password != newPassword {
		return fiber.NewError(400, "Wrong credentials")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte("hi"))

	if err != nil {
		return fiber.NewError(400, "Something went wrong")
	}

	cookie := fiber.Cookie{
		Name:     "JWT",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"Message": "Success",
	})
}
