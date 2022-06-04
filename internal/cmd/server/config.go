package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

const (
	viewDir = "./views"
	viewExt = ".html"
)

func getConfigs() fiber.Config {
	return fiber.Config{
		Views: html.New(viewDir, viewExt),
	}
}
