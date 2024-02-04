package handlers

import (
	"log"
	"net/http"

	converter "github.com/ddefrancesco/scoperunner_server/commands/converters"
	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	handler "github.com/ddefrancesco/scoperunner_server/handlers/commons"
	commons "github.com/ddefrancesco/scoperunner_server/models/commons"
	"github.com/ddefrancesco/scoperunner_server/scopeparser"
	"github.com/gorilla/mux"
)

func InfoCommandHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("InfoCommandHandler::Init -> eseguito")
	vars := mux.Vars(r)
	var infoParam = vars["infos"]
	log.Printf("InfoCommandHandler::Command::PathParam -> %s ###", infoParam)
	//GET Requests
	tMap, err := converter.RequestParamsToInfoArray(infoParam)
	if err != nil {
		appErr := &commons.ScopeErr{
			Err:            http.StatusBadRequest,
			ErrDescription: "Error parsing command: Opzione non valida",
			ScopeFunction:  "Get Info",
			Cmd:            infoParam,
		}
		handler.JSONError(w, appErr, http.StatusBadRequest)
		return
	}
	var infoCommandArray []scopeparser.InfoCommand
	var scopeResponses []interfaces.ScopeResponse
	for _, v := range tMap {
		infoCommand := *scopeparser.NewInfoCommand(v)
		infoCommandArray = append(infoCommandArray, infoCommand)
	}
	serialDevice := handler.GetScopeClient()
	for _, v := range infoCommandArray {

		ic, err := v.ParseMap()
		log.Printf("InfoCommandHandler::Command::Info -> %s ###", v.StringValue())
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

		scopeResponse := serialDevice.FetchQuery(v.StringValue())
		if scopeResponse.Err != nil {
			log.Fatal("Error executing command: porta seriale non trovata")
		}
		log.Printf("InfoCommandHandler::Response %s, \n\t  %s", scopeResponse.Response, scopeResponse.ExecCmd)
		scopeResponses = append(scopeResponses, scopeResponse)
	}

	log.Println("InfoCommandHandler::End -> eseguito")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(handler.SendResponses(r, scopeResponses))
}
