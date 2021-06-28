package handler

import (
	"log"
	"net/http"

	"github.com/xiaoming857/chatroom/api/config"
	"github.com/xiaoming857/chatroom/api/model"
)

// HandleConnections
// HTTP route handler runs similarly to goroutines. This allows the handler to handle multiple requests
// independently from each other.
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP request to websocket protocol
	if ws, err := config.Upgrader.Upgrade(w, r, nil); err != nil {
		log.Fatalln(err)
	} else {
		defer ws.Close()          // Close connection (only when exit this function)
		config.Clients[ws] = true // Register new client
		log.Println(`A User has successfully connected!`)

		// Loop and wait for messages from a client
		for {
			var msg model.ChatMessage

			// Wait for incoming message and read JSON into ChatMessage struct
			if err := ws.ReadJSON(&msg); err != nil {
				log.Println(`Failed to read message: ` + err.Error())
				delete(config.Clients, ws)
				break
			}
			config.Broadcaster <- msg // Send message to HandleMessages for broadcasting
		}
	}
}

// HandleMessages
// Listening to incoming messages and broadcast them to all clients
func HandleMessages() {
	for {
		msg := <-config.Broadcaster // Listen to stream for incoming messages
		log.Println(`Retrieved message from ` + msg.Username + `...`)

		// Broadcast messages to all clients
		for client := range config.Clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Println(`Error: ` + err.Error())
				client.Close()
				delete(config.Clients, client)
			}
		}
	}
}
