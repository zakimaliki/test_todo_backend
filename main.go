package main

import (
	"log"
	"todo-api/pkg/config"
	"todo-api/pkg/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()

	// Inisialisasi koneksi ke Oracle
	config.InitializeDB()

	// Mengatur routing
	routes.SetupRoutes(app)

	log.Println("Server is running on port 3000")
	app.Listen(":3000")
}
