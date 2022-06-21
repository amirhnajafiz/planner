package server

import (
	"database/sql"
	"os"

	"github.com/amirhnajafiz/planner/internal/handler"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	Logger *zap.Logger
	Db     *sql.DB
}

func (s Server) New() {
	// creating a new fiber
	app := fiber.New(getConfigs())

	// defining a new handler
	h := handler.Handler{
		Db:     s.Db,
		Logger: s.Logger.Named("handler"),
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
	s.Logger.Error(app.Listen(":" + port).Error())
}
