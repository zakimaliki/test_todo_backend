package routes

import (
	"todo-api/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/tasks", controllers.CreateTask)
	api.Get("/tasks", controllers.GetAllTasks)
	api.Get("/tasks/:id", controllers.GetTaskByID)
	api.Put("/tasks/:id", controllers.UpdateTask)
	api.Delete("/tasks/:id", controllers.DeleteTask)
}
