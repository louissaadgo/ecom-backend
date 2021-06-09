package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ecom-backend/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", register)
	app.Post("api/login", login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("api/user", user)
	app.Post("api/logout", logout)

	app.Get("api/users", users)
	app.Post("api/users", createUser)
	app.Get("api/users/:id", getUser)
	app.Put("api/users/:id", updateUser)
	app.Delete("api/users/:id", deleteUser)

	app.Get("api/roles", roles)
	app.Post("api/roles", createRole)
	app.Get("api/roles/:id", getRole)
	app.Put("api/roles/:id", updateRole)
	app.Delete("api/roles/:id", deleteRole)

	app.Get("api/permissions", permissions)
}
