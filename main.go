package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/louissaadgo/ecom-backend/database"
	"github.com/louissaadgo/ecom-backend/routes"
)

const port = ":3000"

func main() {

	app := fiber.New()

	database.Connect()

	defer database.DB.Close()

	routes.Setup(app)

	app.Listen(port)
}
