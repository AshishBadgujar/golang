package main

import (
	"go-social/internal/env"

	"log"
)

func main() {
	// Load environment variables from .env file
	if err := env.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	app := &api{
		config: config{
			addr: env.GetString("ADDR", ":8080"),
		},
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
