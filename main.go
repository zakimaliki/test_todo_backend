package main

import (
	"log"
	"todo-api/pkg/config"
	"todo-api/pkg/routes"
	"todo-api/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()

	config.InitRedis()
	// Inisialisasi koneksi ke Oracle
	config.InitializeDB()
	utils.Migration()

	// Mengatur routing
	routes.SetupRoutes(app)

	log.Println("Server is running on port 3000")
	app.Listen(":3000")
}
