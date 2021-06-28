package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xiaoming857/chatroom/api/config"
	"github.com/xiaoming857/chatroom/api/route"
)

func main() {
	port := config.GetPort()

	r := mux.NewRouter()
	route.Setup(r)

	log.Println(`Serve starting at localhost:` + port)
	log.Fatalln(http.ListenAndServe(`:`+port, r))
}
