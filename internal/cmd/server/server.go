package server

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func New() {
	app := fiber.New()

	port := os.Getenv("port")
	if port == "" {
		port = "8080"
	}

	log.Fatalln(app.Listen(":" + port))
}
