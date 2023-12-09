package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ddefrancesco/scoperunner_server/etxclient"
	handler "github.com/ddefrancesco/scoperunner_server/handlers/commons"
	commons "github.com/ddefrancesco/scoperunner_server/models/commons"
)

func AckCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AckCommandHandler::Init -> eseguito")
	// vars := mux.Vars(r)
	// amode := vars["mode"]
	var req commons.ScopeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//GET Request

	command_string := "ACK <0x06>"

	etx := etxclient.NewClient()
	scopeResp := etx.ExecReturnData(command_string)
	if scopeResp.Err != nil {
		log.Fatal("Error executing command: porta seriale non trovata")
	}
	log.Println("AckCommandHandler::End -> eseguito")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(handler.SendResponse(scopeResp))
}
