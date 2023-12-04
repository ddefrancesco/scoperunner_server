package main

import (
	"log"
	"net/http"

	"github.com/ddefrancesco/scoperunner_server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server::Init -> eseguito")
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/align", handlers.AlignCommandHandler)
	log.Println("Server::NewRoute -> registrata")
	log.Println("Server::Bind a porta 8000 -> eseguito")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
