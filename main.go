package main

import (
	"fmt"
	"go-smartcerti/database"
	"go-smartcerti/initializers"
	"go-smartcerti/migrations"
	"go-smartcerti/routes"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	// Load environment variables
	initializers.LoadEnvVariables()
	//initialize database connection
	database.DatabaseInit()
	//initialize database migration
	migrations.Migration()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello!",
		})
	})

	//initialize routes
	routes.RouteInit(app)
	fmt.Println("Server running on port", os.Getenv("PORT"))
	app.Listen(os.Getenv("PORT"))
}
