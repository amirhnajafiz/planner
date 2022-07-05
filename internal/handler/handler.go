package handler

import (
	"database/sql"
	"errors"

	"github.com/amirhnajafiz/planner/internal/debug"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Handler for endpoints
type Handler struct {
	Db     *sql.DB
	Logger *zap.Logger
}

// postgresQL queries
const (
	selectAllQuery  = "SELECT item FROM todos"
	insertItemQuery = "INSERT into todos(item) VALUES ($1)"
	updateItemQuery = "UPDATE todos SET item=$1 WHERE item=$2"
	deleteItemQuery = "DELETE FROM todos WHERE item=$1"
)

// Register our handler endpoints routes
func (h Handler) Register(app *fiber.App) {
	// creating our endpoints
	app.Get("/", h.homePage)            // home
	app.Post("/", h.postHandler)        // post handler
	app.Put("/update", h.putHandler)    // update post handler
	app.Delete("/delete", h.delHandler) // delete post handler
}

func (h Handler) homePage(c *fiber.Ctx) error {
	var (
		// query extracting variables
		res   string
		todos []string
	)

	// executing our query
	rows, err := h.Db.Query(selectAllQuery)

	defer rows.Close()

	if err != nil {
		h.Logger.Error(debug.DatabaseError, zap.Error(err))

		return errors.New(debug.DatabaseError)
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

	newTodo := todo{}

	// parsing the body
	if err := c.BodyParser(&newTodo); err != nil {
		h.Logger.Error(debug.ParsingError, zap.Error(err))

		return err
	}

	// save the new item
	if newTodo.Item != "" {
		_, err := h.Db.Exec(insertItemQuery, newTodo.Item)
		if err != nil {
			h.Logger.Error(debug.SavingError, zap.Error(err))
		}
	}

	return c.Redirect("/")
}

func (h Handler) putHandler(c *fiber.Ctx) error {
	// update database
	if _, err := h.Db.Exec(updateItemQuery, c.Query("newitem"), c.Query("olditem")); err != nil {
		h.Logger.Error(debug.UpdateError, zap.Error(err))

		return err
	}

	return c.Redirect("/")
}

func (h Handler) delHandler(c *fiber.Ctx) error {
	// remove from database
	if _, err := h.Db.Exec(deleteItemQuery, c.Query("item")); err != nil {
		h.Logger.Error(debug.DeleteError, zap.Error(err))

		return err
	}

	return c.SendString("deleted")
}
