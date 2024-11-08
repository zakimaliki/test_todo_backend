package routes

import (
	"todo-api/pkg/controllers"
	"todo-api/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/tasks", middlewares.JwtMiddleware(), controllers.CreateTask)
	app.Get("/tasks", middlewares.JwtMiddleware(), controllers.GetAllTasks)
	app.Get("/tasks/:id", middlewares.JwtMiddleware(), controllers.GetTaskByID)
	app.Put("/tasks/:id", middlewares.JwtMiddleware(), controllers.UpdateTask)
	app.Delete("/tasks/:id", middlewares.JwtMiddleware(), controllers.DeleteTask)
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)
}
