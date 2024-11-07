package main

import (
	"log"
	"todo-api/pkg/config"
	"todo-api/pkg/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Inisialisasi koneksi ke Oracle
	config.InitializeDB()

	// Mengatur routing
	routes.SetupRoutes(app)

	log.Println("Server is running on port 3000")
	app.Listen(":3000")
}
