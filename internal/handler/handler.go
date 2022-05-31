package handler

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {
	// creating our endpoints
	app.Get("/", index)               // home
	app.Post("/", postHandler)        // post handler
	app.Put("/update", putHandler)    // update post handler
	app.Delete("/delete", delHandler) // delete post handler
}

func index(c *fiber.Ctx) error {
	return c.SendString("hello world")
}

func postHandler(c *fiber.Ctx) error {
	return c.SendString("hello world")
}

func putHandler(c *fiber.Ctx) error {
	return c.SendString("hello world")
}

func delHandler(c *fiber.Ctx) error {
	return c.SendString("hello world")
}
