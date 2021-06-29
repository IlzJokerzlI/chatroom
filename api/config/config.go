package config

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/xiaoming857/chatroom/api/model"
)

// Global variables
var (
	Clients     = make(map[*websocket.Conn]bool) // Map of clients (uses maps to directly map the key object)
	Broadcaster = make(chan model.ChatMessage)   // A channel between HandleConnections and HandleMessages
	Upgrader    = websocket.Upgrader{            // Upgrade HTTP request to WebSocket protocol
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// GetPort
// Load environment variables (PORT)
func GetPort() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(`Error loading .env file`)
	}
	port := os.Getenv(`PORT`)
	return port
}
