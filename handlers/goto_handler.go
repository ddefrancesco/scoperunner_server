package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	handler "github.com/ddefrancesco/scoperunner_server/handlers/commons"
	commons "github.com/ddefrancesco/scoperunner_server/models/commons"
	scopeparser "github.com/ddefrancesco/scoperunner_server/scopeparser"
)

func GotoCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GotoCommandHandler::Init -> eseguito")
	// vars := mux.Vars(r)
	//amode := vars["mode"]

	var moves commons.ScopeSetRequest
	err := json.NewDecoder(r.Body).Decode(&moves)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	moveRequest := scopeparser.NewMoveRequest(moves.Body)

	ac, err := moveRequest.ParseMap()
	if err != nil {
		appErr := &commons.ScopeErr{
			Err:            http.StatusBadRequest,
			ErrDescription: "Error parsing command: Opzione non valida",
			ScopeFunction:  "Move",
			Cmd:            ac,
		}
		handler.JSONError(w, appErr, http.StatusBadRequest)
		return
	}

	serialDevice := handler.GetScopeClient()
	//command_string := alignCmd.ParseCommand()
	var scopeResponseArray []interfaces.ETXResponse
	for k, v := range ac {
		command_string := ":" + k + v + "#"
		log.Printf("GotoCommandHandler::Command::Info -> %s ###", command_string)
		scopeResp := serialDevice.ExecCommand(command_string)
		if scopeResp.Err != nil {
			log.Fatal("Error executing command: porta seriale non trovata")
		}
		scopeResponseArray = append(scopeResponseArray, scopeResp)
	}
	log.Println("GotoCommandHandler::End -> eseguito")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	for _, v := range scopeResponseArray {
		w.Write(handler.SendResponse(r, v))
	}

}
