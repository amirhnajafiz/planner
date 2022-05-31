package server

import (
	"log"
	"os"

	"github.com/amirhnajafiz/planner/internal/db"
	"github.com/amirhnajafiz/planner/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func New() {
	// creating a new fiber
	app := fiber.New()

	// creating a new database connection
	d, err := db.NewConnection()
	if err != nil {
		panic(any(err))
	}

	// defining a new handler
	h := handler.Handler{
		Db: d,
	}

	// registering our application
	h.Register(app)

	// getting the port from env variables
	port := os.Getenv("port")
	if port == "" {
		port = "8080" // default port is 8080
	}

	// starting our server
	log.Fatalln(app.Listen(":" + port))
}
