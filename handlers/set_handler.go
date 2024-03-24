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

func SetCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SetCommandHandler::Init -> eseguito")
	// vars := mux.Vars(r)
	//amode := vars["mode"]

	var settings commons.ScopeSetRequest
	err := json.NewDecoder(r.Body).Decode(&settings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	setRequest := scopeparser.NewSetRequest(settings.Body)

	ac, err := setRequest.ParseMap()
	if err != nil {
		appErr := &commons.ScopeErr{
			Err:            http.StatusBadRequest,
			ErrDescription: "Error parsing command: Opzione non valida",
			ScopeFunction:  "Set",
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
		log.Printf("SetCommandHandler::Command::Info -> %s ###", command_string)
		scopeResp := serialDevice.ExecCommand(command_string)
		if scopeResp.Err != nil {
			log.Fatal("Error executing command: porta seriale non trovata")
		}
		scopeResponseArray = append(scopeResponseArray, scopeResp)
	}
	log.Println("SetCommandHandler::End -> eseguito")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	for _, v := range scopeResponseArray {
		w.Write(handler.SendResponse(r, v))
	}

}
