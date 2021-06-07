package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	app.Post("/api/register", register)
	app.Post("api/login", login)
	app.Get("api/user", user)
}
