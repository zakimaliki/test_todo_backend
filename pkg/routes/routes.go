package routes

import (
	"todo-api/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/tasks", controllers.CreateTask)
	app.Get("/tasks", controllers.GetAllTasks)
	app.Get("/tasks/:id", controllers.GetTaskByID)
	app.Put("/tasks/:id", controllers.UpdateTask)
	app.Delete("/tasks/:id", controllers.DeleteTask)
}
