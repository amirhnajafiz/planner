package server

import (
	"log"
	"os"

	"github.com/amirhnajafiz/planner/internal/db"
	"github.com/amirhnajafiz/planner/internal/handler"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func New(l *zap.Logger) {
	// creating a new fiber
	app := fiber.New(getConfigs())

	// creating a new database connection
	d, err := db.NewConnection()
	if err != nil {
		panic(any(err))
	}

	// defining a new handler
	h := handler.Handler{
		Db:     d,
		Logger: l,
	}

	// registering our application
	h.Register(app)

	// getting the port from env variables
	port := os.Getenv("port")
	if port == "" {
		port = "8080" // default port is 8080
	}

	// adding assets files
	app.Static("/", "./public")

	// starting our server
	log.Fatalln(app.Listen(":" + port))
}
