package main

import (
	"fiber/database"
	"fiber/lead"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/lead", lead.GetLead)
	app.Get("/api/lead/:id", lead.GetLead)
	app.Post("/api/lead", lead.NewLead)
	app.Delete("/api/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}
func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
	defer database.DBConn.Close()

}
