package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ddefrancesco/scoperunner_server/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("scope-server-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error reading in config: ", err)
	}
	log.Println("Server::Init -> eseguito")
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/align", handlers.AlignCommandHandler).Methods("POST")
	log.Println("Server::NewRoute /align -> registrata")
	r.HandleFunc("/ack", handlers.AckCommandHandler).Methods("GET")
	log.Println("Server::NewRoute /ack -> registrata")
	r.HandleFunc("/info/{item}", handlers.InfoCommandHandler).Methods("GET")
	log.Println("Server::NewRoute /info -> registrata")
	log.Println("Server::Bind a porta 8000 -> eseguito")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
