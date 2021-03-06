package routes

import (
	"github.com/bsgitt/golang-fiber-basic-todo-app/controllers" // replace
	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("/", controllers.GetTodos)
	route.Post("/", controllers.CreateTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Get("/:id", controllers.GetTodo)
}
