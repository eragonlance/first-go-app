package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := setup()
	log.Fatal(app.Listen(":8000", fiber.ListenConfig{
		DisableStartupMessage: IsDeployed(),
	}))
}

func getVersion(c fiber.Ctx) error {
	if !IsDeployed() {
		return c.SendString("Applicable for deployed only.")
	}

	ref := os.Getenv("GIT_REF")
	sha := os.Getenv("GIT_SHA")
	return c.SendString(fmt.Sprintf("%s\n%s", ref, sha))
}

// Set up app.
//
// Pass in "test" to exclude unnecessary stuff when implementing tests
func setup(args ...any) *fiber.App {
	test := false
	for _, arg := range args {
		test = arg == "test"
	}

	app := fiber.New()
	if !IsDeployed() && !test {
		app.Use(logger.New())
	}

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/version", getVersion)
	return app
}

// Whether the environment is deployed or local
func IsDeployed() bool {
	return os.Getenv("DEPLOYED") == "1"
}
