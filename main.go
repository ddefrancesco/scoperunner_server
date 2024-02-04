package main

import (
	"log"
	"net/http"

	"github.com/ddefrancesco/scoperunner_server/handlers"
	"github.com/gorilla/mux"

	configuration "github.com/ddefrancesco/scoperunner_server/configurations"
)

func main() {
	err := configuration.InitConfig()
	if err != nil {
		panic(err)
	}
	log.Println("Server::Init -> eseguito")
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/align", handlers.AlignCommandHandler).Methods("POST")
	log.Println("Server::NewRoute /align -> registrata")
	r.HandleFunc("/ack", handlers.AckCommandHandler).Methods("GET")
	log.Println("Server::NewRoute /ack -> registrata")
	r.HandleFunc("/info/{infos}", handlers.InfoCommandHandler).Methods("GET")
	log.Println("Server::NewRoute /info -> registrata")
	log.Println("Server::Bind a porta 8000 -> eseguito")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
