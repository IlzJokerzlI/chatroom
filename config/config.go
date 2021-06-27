package config

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/xiaoming857/chatroom/model"
)

var (
	Clients     = make(map[*websocket.Conn]bool)
	Broadcaster = make(chan model.ChatMessage)
	Upgrader    = websocket.Upgrader{
		CheckOrigin:     func(r *http.Request) bool { return true },
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func GetPort() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(`Error loading .env file`)
	}
	port := os.Getenv(`PORT`)
	return port
}
