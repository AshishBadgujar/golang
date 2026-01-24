package routes

import (
	"todo-app/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, todoHandler *handlers.TodoHandler) {
	api := app.Group("/api")

	todos := api.Group("/todos")
	todos.Post("/", todoHandler.Create)
	todos.Get("/", todoHandler.GetAll)
	todos.Get("/:id", todoHandler.GetByID)
	todos.Patch("/:id", todoHandler.Update)
	todos.Delete("/:id", todoHandler.Delete)
}
