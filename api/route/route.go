package route

import (
	"github.com/gorilla/mux"
	"github.com/xiaoming857/chatroom/api/handler"
)

// Setup
// Setup route
func Setup(r *mux.Router) {
	r.HandleFunc(`/`, handler.Home)                // Home
	r.HandleFunc(`/ws`, handler.HandleConnections) // Handle WebSocket
	go handler.HandleMessages()                    // Listening to incoming messages
}
