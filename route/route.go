package route

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xiaoming857/chatroom/handler"
)

// Setup
func Setup(r *mux.Router) {
	r.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `Hello World!`)
	})
	r.HandleFunc(`/ws`, handler.HandleConnections)
	go handler.HandleMessages() // Listening for incoming messages
}
