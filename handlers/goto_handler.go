package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ddefrancesco/scoperunner_server/errors"
	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	handler "github.com/ddefrancesco/scoperunner_server/handlers/commons"
	commons "github.com/ddefrancesco/scoperunner_server/models/commons"
	scopeparser "github.com/ddefrancesco/scoperunner_server/scopeparser"
)

func GotoCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GotoCommandHandler::Init -> eseguito")

	/** {
		 	"body": {
	             "goto": "M31",
		          },
		 }*/
	var settings commons.ScopeInitRequest
	err := json.NewDecoder(r.Body).Decode(&settings.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("GotoCommandHandler::Settings::Body -> %v ###", settings)

	gotoJsonMap := settings.Body["body"]
	gotoRequest := scopeparser.NewGotoRequest(gotoJsonMap)

	gotoRADecCmd, err := gotoRequest.FindDeepSpaceObjectCommand()
	switch e := err.(type) {
	case *errors.ObjectNotFoundInCatalogError:
		// Handle the specific error
		log.Printf("Object not found in catalog: %s", e.Message)
		//Ricerca oggetto stellare o pianeta

		http.Error(w, fmt.Sprintf("Error: %s", e.Message), http.StatusNotFound)
	default:
		// Handle other types of errors
		log.Printf("Unexpected error: %v", err)
		appErr := &commons.ScopeErr{
			Err:            http.StatusBadRequest,
			ErrDescription: "Error parsing command: Parametro non valida",
			ScopeFunction:  "Initialize",
			Cmd:            gotoRADecCmd,
		}
		handler.JSONError(w, appErr, http.StatusBadRequest)
		// You can send a generic error response to the client here
		// For example:
		// http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	// ac, err := initRequest.ParseMap()
	if err != nil {

		return
	}

	// Send command to serial device
	//handler := &handlers.Handler{}
	// Get the scope client
	// Get the serial device
	serialDevice := handler.GetScopeClient()
	//command_string := alignCmd.ParseCommand()
	var scopeResponseArray []interfaces.ETXResponse

	log.Printf("GotoCommandHandler::Command::Info -> %s ###", gotoRADecCmd)
	scopeResp := serialDevice.ExecCommand(gotoRADecCmd)
	if scopeResp.Err != nil {
		log.Fatal("Error executing command: porta seriale non trovata")
	}

	scopeResponseArray = append(scopeResponseArray, scopeResp)
	checkRADecCmd, err := gotoRequest.CheckGotoRADecCommand()
	// ac, err := initRequest.ParseMap()
	scopeResp = serialDevice.ExecCommand(checkRADecCmd)
	if scopeResp.Err != nil {
		log.Fatal("Error executing command: porta seriale non trovata")
	}
	scopeResponseArray = append(scopeResponseArray, scopeResp)

	if strings.ContainsAny(string(scopeResp.Response), "12") {
		log.Println("GotoCommandHandler::Error -> RA/Dec Sotto l'Orizzonte/Higher")
		appErr := &commons.ScopeErr{
			Err:            http.StatusBadRequest,
			ErrDescription: "Error executing command: RA/Dec Sotto l'Orizzonte/Higher",
			ScopeFunction:  "Goto",
			Cmd:            checkRADecCmd,
		}
		handler.JSONError(w, appErr, http.StatusBadRequest)
		return

	}
	gotoCmd, _ := gotoRequest.SetGotoCommand()
	scopeResp = serialDevice.ExecCommand(gotoCmd)
	if scopeResp.Err != nil {
		log.Fatal("Error executing command: porta seriale non trovata")
	}
	scopeResponseArray = append(scopeResponseArray, scopeResp)
	log.Println("GotoCommandHandler::End -> eseguito")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	for _, v := range scopeResponseArray {
		w.Write(handler.SendResponse(r, v))
	}

}
