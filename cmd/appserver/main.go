package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := setup()
	log.Fatal(app.Listen(":8000"))
}

func getVersion(c fiber.Ctx) error {
	if os.Getenv("DEPLOYED") != "1" {
		return c.SendString("Applicable for deployed only.")
	}

	ref := os.Getenv("GIT_REF")
	sha := os.Getenv("GIT_SHA")
	return c.SendString(fmt.Sprintf("%s\n%s", ref, sha))
}

func setup() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/version", getVersion)
	return app
}
