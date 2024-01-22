package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

const dbName = "fiber-hrms"
const mongoURI = "mongodb://username:password@localhost:27017"

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

func Connect() error {
	clientOption := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	collection = client.Database(dbName).Collection("employees")
	return nil
}

func main() {

	err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {
		query := bson.D{{}}
		cursor, err := collection.Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var employees []Employee = make([]Employee, 0)
		err = cursor.All(c.Context(), &employees)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(employees)
	})

	app.Post("/employee", func(c *fiber.Ctx) error {
		employee := new(Employee)
		err := c.BodyParser(employee)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		employee.ID = ""
		insertionResult, err := collection.InsertOne(c.Context(), employee)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.M{"_id": insertionResult.InsertedID}
		createdRecord := collection.FindOne(c.Context(), filter)
		createdEmployee := Employee{}
		createdRecord.Decode(&createdEmployee)
		return c.Status(201).JSON(createdEmployee)
	})

	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		employee := new(Employee)
		err = c.BodyParser(employee)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		query := bson.M{"_id": id} //convert Json to bson
		update := bson.M{
			"$set": bson.M{
				"name":   employee.Name,
				"age":    employee.Age,
				"salary": employee.Salary,
			}} //convert Json to bson

		_, err = collection.UpdateOne(c.Context(), query, update)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		employee.ID = idParam

		return c.Status(200).JSON(employee)

	})

	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		query := bson.M{"_id": id}
		_, err = collection.DeleteOne(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(200).SendString(idParam)
	})

	log.Fatal(app.Listen(":3000"))
}
