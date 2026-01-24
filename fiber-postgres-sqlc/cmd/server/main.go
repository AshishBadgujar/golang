package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"todo-app/internal/config"
	"todo-app/internal/database"
	"todo-app/internal/db/sqlc"
	"todo-app/internal/handlers"
	"todo-app/internal/routes"
	"todo-app/internal/services"
)

func main() {
	// ------------------------------------------------------------------
	// Load env (safe in Docker & local)
	// ------------------------------------------------------------------
	_ = godotenv.Load()

	cfg := config.Load()

	// ------------------------------------------------------------------
	// Database (pgx pool)
	// ------------------------------------------------------------------
	pool := database.Connect(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	defer pool.Close()

	queries := sqlc.New(pool)

	// Optional: verify DB connectivity on startup
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("database ping failed: %v", err)
	}

	// ------------------------------------------------------------------
	// Fiber app
	// ------------------------------------------------------------------
	app := fiber.New(fiber.Config{
		AppName: "Todo API",
	})

	// ------------------------------------------------------------------
	// Dependency wiring
	// ------------------------------------------------------------------
	todoService := services.NewTodoService(queries)
	todoHandler := handlers.NewTodoHandler(todoService)

	routes.Register(app, todoHandler)

	// ------------------------------------------------------------------
	// Graceful shutdown
	// ------------------------------------------------------------------
	go func() {
		if err := app.Listen(":" + cfg.AppPort); err != nil {
			log.Printf("fiber stopped: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")

	_, shutdownCancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer shutdownCancel()

	if err := app.Shutdown(); err != nil {
		log.Printf("fiber shutdown error: %v", err)
	}

	// pgx pool will close via defer
	log.Println("server exited cleanly")
}
