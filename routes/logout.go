package routes

import "github.com/gofiber/fiber/v2"

func logout(c *fiber.Ctx) error {
	c.ClearCookie("JWT")
	return c.SendString("Logged out")
}
