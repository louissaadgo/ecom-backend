package routes

import (
	"database/sql"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/database"
)

func user(c *fiber.Ctx) error {

	type Claims struct {
		jwt.StandardClaims
	}

	cookie := c.Cookies("JWT")

	if cookie == "" {
		return fiber.NewError(400, "User Not Authenticated")
	}

	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("hi"), nil
	})

	if err != nil || !token.Valid {
		return fiber.NewError(400, "User Not Authenticated")
	}

	claims := token.Claims.(*Claims)

	query := database.DB.QueryRow(`SELECT firstName FROM users WHERE id = $1;`, claims.Issuer)
	var username string
	if err := query.Scan(&username); err == sql.ErrNoRows {
		return fiber.NewError(400, "Somwthing went wrong")
	}

	return c.JSON(username)
}
