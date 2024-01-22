package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

var (
	clients = make(map[*websocket.Conn]string)
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Chat App Server is running!")
	})

	app.Get("/ws", websocket.New(handleWebSocket))

	// Start the server
	err := app.Listen(":8000")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func handleWebSocket(c *websocket.Conn) {
	// Get username from query parameters
	username := c.Query("username")
	clients[c] = username

	// Send a welcome message to the user
	// welcomeMsg := Message{
	// 	Username: "Server",
	// 	Text:     fmt.Sprintf("Welcome, %s!", username),
	// }
	// c.WriteJSON(welcomeMsg)

	// Broadcast user join message
	broadcast(Message{
		Username: "Server",
		Text:     fmt.Sprintf("%s joined the chat", username),
	}, c)

	// Listen for incoming messages
	for {
		var msg Message
		err := c.ReadJSON(&msg)
		if err != nil {
			delete(clients, c)
			broadcast(Message{
				Username: "Server",
				Text:     fmt.Sprintf("%s left the chat", username),
			}, c)
			break
		}

		broadcast(msg, c)
	}
}

func broadcast(msg Message, sender *websocket.Conn) {
	for client := range clients {
		if client == sender {
			continue
		}

		// If the sender is the server, keep the "Server" username in the message
		if msg.Username == "Server" {
			client.WriteJSON(msg)
		} else {
			client.WriteJSON(msg)
		}

	}
}
