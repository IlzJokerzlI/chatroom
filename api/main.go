package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xiaoming857/chatroom/api/config"
	"github.com/xiaoming857/chatroom/api/route"
)

func main() {
	port := config.GetPort() // Get port from environment

	r := mux.NewRouter() // Use Mux framework
	route.Setup(r)       // Setup route

	log.Println(`Serve starting at localhost:` + port) // Serve
	log.Fatalln(http.ListenAndServe(`:`+port, r))
}
