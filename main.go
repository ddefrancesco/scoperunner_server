package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/ddefrancesco/scoperunner_server/handlers"
	"github.com/dotse/go-health/client"
	"github.com/gorilla/mux"

	configuration "github.com/ddefrancesco/scoperunner_server/configurations"
)

func main() {
	//check if healthcheck
	log.Println("Server::Health -> inizializzazione")
	if len(os.Args) >= 2 && os.Args[1] == "healthcheck" {
		client.CheckHealthCommand()
		log.Println("Server::CheckHealthCommand -> eseguito")
	}
	log.Println("Server::Init -> inizializzazione")
	log.Println("Server::CheckInternetConnection -> eseguito")
	if !CheckInternetConnection() {
		log.Println("Server::CheckInternetConnection -> connessione internet non disponibile")
		return
	}
	err := configuration.InitConfig()
	if err != nil {
		panic(err)
	}
	log.Println("Server::InitConfig -> eseguito")
	err = configuration.InitEnvConfig()
	if err != nil {
		panic(err)
	}
	log.Println("Server::InitEnvConfig -> eseguito")
	log.Println("Server::Init -> eseguito")
	r := mux.NewRouter()
	health := mux.NewRouter()

	health.HandleFunc("/health", handlers.HealthCommandHandler).Methods("GET")
	log.Println("Server::NewRoute@ port 9999 /health -> registrata")
	// Routes consist of a path and a handler function.
	r.HandleFunc("/align", handlers.AlignCommandHandler).Methods("POST")
	log.Println("Server::NewRoute /align -> registrata")
	r.HandleFunc("/ack", handlers.AckCommandHandler).Methods("GET")
	log.Println("Server::NewRoute /ack -> registrata")
	r.HandleFunc("/info/{infos}", handlers.InfoCommandHandler).Methods("GET")
	log.Println("Server::NewRoute /info -> registrata")
	r.HandleFunc("/set", handlers.SetCommandHandler).Methods("POST")
	log.Println("Server::NewRoute /set -> registrata")
	r.HandleFunc("/init", handlers.InitCommandHandler).Methods("POST")
	log.Println("Server::NewRoute /init -> registrata")
	r.HandleFunc("/move", handlers.GotoCommandHandler).Methods("POST")
	log.Println("Server::NewRoute /move -> registrata")
	log.Println("Server::Bind a porta 8000 -> eseguito")

	go http.ListenAndServe(":9999", health)
	log.Println("Server::Bind a porta 9999 -> eseguito")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

}

func CheckInternetConnection() bool {
	timeout := 5 * time.Second
	_, err := net.DialTimeout("tcp", "8.8.8.8:53", timeout)
	return err == nil
}
