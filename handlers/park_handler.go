package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	scopecommand "github.com/ddefrancesco/scoperunner_server/commands"
	handler "github.com/ddefrancesco/scoperunner_server/handlers/commons"
	commons "github.com/ddefrancesco/scoperunner_server/models/commons"
	scopeparser "github.com/ddefrancesco/scoperunner_server/scopeparser"
)

func ParkCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ParkCommandHandler::Init -> eseguito")
	// vars := mux.Vars(r)
	// amode := vars["mode"]

	var apark commons.ScopeSetRequest
	err := json.NewDecoder(r.Body).Decode(&apark)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parking := scopeparser.NewParking(setParking(apark.Body["parking"]))

	ac, err := parking.ParseMap()
	if err != nil {
		appErr := &commons.ScopeErr{
			Err:            http.StatusBadRequest,
			ErrDescription: "Error parsing command: Opzione non valida",
			ScopeFunction:  "Parking",
			Cmd:            apark.Body,
		}
		handler.JSONError(w, appErr, http.StatusBadRequest)
		return
	}

	alignCmd := scopecommand.NewParkingCommand(ac)

	command_string := alignCmd.ParseCommand()

	serialDevice := handler.GetScopeClient()
	scopeResp := serialDevice.ExecCommand(command_string)
	if scopeResp.Err != nil {
		log.Fatal("Error executing command: porta seriale non trovata")
	}
	log.Println("AlignCommandHandler::End -> eseguito")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(handler.SendResponse(r, scopeResp))
}

func setParking(mode string) scopeparser.ParkType {
	switch mode {
	case "park":
		return scopeparser.ParkTypePark
	case "seek":
		return scopeparser.ParkTypeSeek
	case "query":
		return scopeparser.ParkTypeQuery
	default:
		return "error"
	}
}
