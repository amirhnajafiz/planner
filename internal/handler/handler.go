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
		// query extracting variables
		res   string
		todos []string
		// index query
		query = "SELECT * FROM todos"
	)

	// executing our query
	rows, err := h.Db.Query(query)

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	if err != nil {
		log.Print(err)

		_ = c.SendString("[Failed] Error in DB")
	}

	// extracting our query results
	for rows.Next() {
		_ = rows.Scan(&res)

		todos = append(todos, res)
	}

	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}

func (h Handler) postHandler(c *fiber.Ctx) error {
	// type fo request
	type todo struct {
		Item string `json:"item"`
	}

	// our query and new item
	query := "INSERT into todos VALUES ($1)"
	newTodo := todo{}

	// parsing the body
	if err := c.BodyParser(&newTodo); err != nil {
		log.Println(err)

		return c.SendString(err.Error())
	}

	// save the new item
	if newTodo.Item != "" {
		_, err := h.Db.Exec(query, newTodo.Item)
		if err != nil {
			log.Println(err)
		}
	}

	return c.Redirect("/")
}

func (h Handler) putHandler(c *fiber.Ctx) error {
	// query
	query := "UPDATE todos SET item=$1 WHERE item=$s"
	// items
	oldItem := c.Query("olditem")
	newItem := c.Query("newitem")

	// update database
	_, _ = h.Db.Exec(query, newItem, oldItem)

	return c.Redirect("/")
}

func (h Handler) delHandler(c *fiber.Ctx) error {
	// query and item
	query := "DELETE form todos WHERE item=$1"
	todoToDelete := c.Query("item")

	// remove from database
	_, _ = h.Db.Exec(query, todoToDelete)

	return c.SendString("deleted")
}
