package server

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func New() {
	// creating a new fiber
	app := fiber.New()

	// getting the port from env variables
	port := os.Getenv("port")
	if port == "" {
		port = "8080" // default port is 8080
	}

	// starting our server
	log.Fatalln(app.Listen(":" + port))
}
