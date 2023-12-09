package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ddefrancesco/scoperunner_server/etxclient"
	scopeparser "github.com/ddefrancesco/scoperunner_server/scopeparser"

	scopecommand "github.com/ddefrancesco/scoperunner_server/commands"

	commons "github.com/ddefrancesco/scoperunner_server/models/commons"

	handler "github.com/ddefrancesco/scoperunner_server/handlers/commons"
)

func AlignCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AlignCommandHandler::Init -> eseguito")
	// vars := mux.Vars(r)
	// amode := vars["mode"]
	var amode commons.ScopeRequest
	err := json.NewDecoder(r.Body).Decode(&amode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	alignment := scopeparser.NewAlignment(setMode(amode.Body))

	ac, err := alignment.ParseMap()
	if err != nil {
		appErr := &commons.ScopeErr{
			Err:            http.StatusBadRequest,
			ErrDescription: "Error parsing command: Opzione non valida",
			ScopeFunction:  "Align",
			Cmd:            amode.Body,
		}
		handler.JSONError(w, appErr, http.StatusBadRequest)
		return
	}

	alignCmd := scopecommand.NewAlignCommand(ac)

	command_string := alignCmd.ParseCommand()

	etx := etxclient.NewClient()
	scopeResp := etx.ExecReturnNothing(command_string)
	if scopeResp.Err != nil {
		log.Fatal("Error executing command: porta seriale non trovata")
	}
	log.Println("AlignCommandHandler::End -> eseguito")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(handler.SendResponse(scopeResp))
}

func setMode(mode string) scopeparser.AlignMode {
	switch mode {
	case "altaz":
		return scopeparser.AltAz
	case "polar":
		return scopeparser.Polar
	case "land":
		return scopeparser.Land
	default:
		return "error"
	}
}
