package handlers

import (
	"log"
	"net/http"

	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	handler "github.com/ddefrancesco/scoperunner_server/handlers/commons"
)

func AckCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AckCommandHandler::Init -> eseguito")
	// vars := mux.Vars(r)
	// amode := vars["mode"]

	//GET Request

	command_string := "ACK <0x06>"

	serialDevice := handler.GetScopeClient()
	scopeResp := serialDevice.ExecCommand(command_string)
	if scopeResp.Err != nil {
		log.Fatal("Error executing command: porta seriale non trovata")
	}
	scopeResps := []interfaces.ETXResponse{scopeResp}
	log.Printf("AckCommandHandler::Response %s, \n\t  %s", scopeResp.Response, scopeResp.ExecCmd)
	log.Println("AckCommandHandler::End -> eseguito")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(handler.SendResponses(r, scopeResps))
}
