package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/database"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func main() {
	database.ConnectionDb()

	app := fiber.New()

	app.Get("/v1", welcome)

	log.Fatal(app.Listen(":3000"))
}
