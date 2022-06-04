package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Db *sql.DB
}

func (h Handler) Register(app *fiber.App) {
	// creating our endpoints
	app.Get("/", h.index)               // home
	app.Post("/", h.postHandler)        // post handler
	app.Put("/update", h.putHandler)    // update post handler
	app.Delete("/delete", h.delHandler) // delete post handler
}

func (h Handler) index(c *fiber.Ctx) error {
	var (
		res   string
		todos []string
	)

	rows, err := h.Db.Query("SELECT * FROM todos")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)

		c.SendString("[Failed] Error in DB")
	}

	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}

	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}

func (h Handler) postHandler(c *fiber.Ctx) error {
	return c.SendString("hello world")
}

func (h Handler) putHandler(c *fiber.Ctx) error {
	return c.SendString("hello world")
}

func (h Handler) delHandler(c *fiber.Ctx) error {
	return c.SendString("hello world")
}
