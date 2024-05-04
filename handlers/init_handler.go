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

func InitCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("InitCommandHandler::Init -> eseguito")

	var settings commons.ScopeInitRequest
	err := json.NewDecoder(r.Body).Decode(&settings.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("InitCommandHandler::Settings::Info -> %v ###", settings)

	addressJsonMap := settings.Body["body"]
	addressJson := commons.RequestAddress{
		Address: addressJsonMap["address"],
	}
	log.Printf("InitCommandHandler::Address::Info -> %v ###", addressJson)

	initRequest := scopeparser.NewInitRequest(addressJson)
	initCmd, err := initRequest.SetInitializeCommand()
	// ac, err := initRequest.ParseMap()
	if err != nil {
		appErr := &commons.ScopeErr{
			Err:            http.StatusBadRequest,
			ErrDescription: "Error parsing command: Parametro non valida",
			ScopeFunction:  "Initialize",
			Cmd:            initCmd,
		}
		handler.JSONError(w, appErr, http.StatusBadRequest)
		return
	}

	serialDevice := handler.GetScopeClient()
	//command_string := alignCmd.ParseCommand()
	var scopeResponseArray []interfaces.ETXResponse
	//for k, v := range ac {
	//command_string := ":" + k + v + "#"
	log.Printf("InitCommandHandler::Command::Info -> %s ###", initCmd)
	scopeResp := serialDevice.ExecCommand(initCmd)
	if scopeResp.Err != nil {
		log.Fatal("Error executing command: porta seriale non trovata")
	}
	scopeResponseArray = append(scopeResponseArray, scopeResp)
	//}
	log.Println("InitCommandHandler::End -> eseguito")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	for _, v := range scopeResponseArray {
		w.Write(handler.SendResponse(r, v))
	}

}
