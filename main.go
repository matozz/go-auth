package main

import (
	"demo/database"
	"demo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		// able to send the cookie back
		AllowCredentials: true,
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
