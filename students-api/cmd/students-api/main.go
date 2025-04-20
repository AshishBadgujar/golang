package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ashishbadgujar/golang/students-api/internal/config"
	"github.com/ashishbadgujar/golang/students-api/internal/http/handlers/student"
	"github.com/ashishbadgujar/golang/students-api/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	db, err := sqlite.New(cfg)
	if err != nil {
		log.Fatalf("Error creating storage: %s", err.Error())
	}
	slog.Info("Storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(db))
	router.HandleFunc("GET /api/students/{id}", student.GetById(db))
	router.HandleFunc("GET /api/students", student.GetAll(db))
	router.HandleFunc("PUT /api/students/{id}", student.Update(db))
	router.HandleFunc("DELETE /api/students/{id}", student.Delete(db))

	server := http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: router,
	}

	slog.Info("Server started on: ", slog.String("addr", cfg.HttpServer.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error starting server: %s", err.Error())

		}
	}()

	<-done

	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Error shutting down server", slog.String("error", err.Error()))
	}

	slog.Info("Server stopped")

}
