package handlers

import (
	"log"
	"net/http"

	handler "github.com/ddefrancesco/scoperunner_server/handlers/commons"
	commons "github.com/ddefrancesco/scoperunner_server/models/commons"
	"github.com/ddefrancesco/scoperunner_server/scopeparser"
	"github.com/gorilla/mux"
)

func InfoCommandHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("InfoCommandHandler::Init -> eseguito")
	vars := mux.Vars(r)
	var infoParam = vars["item"]
	log.Printf("InfoCommandHandler::Command::PathParam -> %s ###", infoParam)
	//GET Requests
	info := scopeparser.NewInfoCommand(scopeparser.Info(infoParam))

	ic, err := info.ParseMap()
	log.Printf("InfoCommandHandler::Command::Info -> %s ###", info.StringValue())
	if err != nil {
		appErr := &commons.ScopeErr{
			Err:            http.StatusBadRequest,
			ErrDescription: "Error parsing command: Opzione non valida",
			ScopeFunction:  "Get Info",
			Cmd:            ic,
		}
		handler.JSONError(w, appErr, http.StatusBadRequest)
		return
	}

	serialDevice := handler.GetScopeClient()

	scopeResp := serialDevice.FetchQuery(info.StringValue())
	if scopeResp.Err != nil {
		log.Fatal("Error executing command: porta seriale non trovata")
	}
	log.Printf("InfoCommandHandler::Response %s, \n\t  %s", scopeResp.Response, scopeResp.ExecCmd)
	log.Println("InfoCommandHandler::End -> eseguito")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(handler.SendResponse(r, scopeResp))
}
