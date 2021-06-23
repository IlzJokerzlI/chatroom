package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xiaoming857/chatroom/config"
)

func main() {
	port := config.GetPort()

	r := mux.NewRouter()
	r.HandleFunc(`/`, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `Hello World!`)
	})
	log.Println(`Serve starting at localhost:` + port)
	log.Fatalln(http.ListenAndServe(`:`+port, r))
}
