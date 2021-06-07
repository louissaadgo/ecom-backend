package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func ParseJwt(cookie string) (string, error) {

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("hi"), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("JWT")
	if _, err := ParseJwt(cookie); err != nil {
		return fiber.NewError(400, "Not Authenticated")
	}
	return c.Next()
}
