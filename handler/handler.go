package handler

import (
	"log"
	"net/http"

	"github.com/xiaoming857/chatroom/config"
	"github.com/xiaoming857/chatroom/model"
)

// HandleConnections
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	if ws, err := config.Upgrader.Upgrade(w, r, nil); err != nil {
		log.Fatalln(err)
	} else {
		defer ws.Close()
		config.Clients[ws] = true

		for {
			var msg model.ChatMessage
			if err := ws.ReadJSON(&msg); err != nil {
				log.Println(`Failed to read message: ` + err.Error())
				delete(config.Clients, ws)
				break
			}
			config.Broadcaster <- msg
		}
	}
}

// HandleMessages
// Listening to incoming messages and broadcast them to all clients
func HandleMessages() {
	for {
		msg := <-config.Broadcaster // Listen to stream
		log.Println(`Retrieved message from ` + msg.Username)

		// Send messages to all clients
		for client := range config.Clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Println(`Error: ` + err.Error())
				client.Close()
				delete(config.Clients, client)
			}
		}
	}
}
