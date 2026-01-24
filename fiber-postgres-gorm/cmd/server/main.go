package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"todo-app/internal/config"
	"todo-app/internal/database"
	"todo-app/internal/handlers"
	"todo-app/internal/models"
	"todo-app/internal/routes"
	"todo-app/internal/services"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	db := database.Connect(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	// Auto migrate
	db.AutoMigrate(&models.Todo{})

	app := fiber.New()

	todoService := services.NewTodoService(db)
	todoHandler := handlers.NewTodoHandler(todoService)

	routes.Register(app, todoHandler)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
