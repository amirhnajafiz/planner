package handler

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	// creating our endpoints
	app.Get("/")
	app.Post("/")
	app.Put("/update")
	app.Delete("/delete")
}
