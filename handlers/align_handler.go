package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ddefrancesco/scoperunner_server/etxclient"
	scopeparser "github.com/ddefrancesco/scoperunner_server/scopeparser"

	scopecommand "github.com/ddefrancesco/scoperunner_server/commands"
	"github.com/gorilla/mux"
)

type scope_response struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func sendResponse(sr etxclient.ScopeResponse) []byte {
	log.Println("sendResponse::Init -> eseguito")

	scoperesponse := scope_response{Code: http.StatusOK, Description: string(sr.Response)}

	jsonResponse, jsonError := json.Marshal(scoperesponse)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	fmt.Println(string(jsonResponse))

	return jsonResponse
}

func AlignCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AlignCommandHandler::Init -> eseguito")
	vars := mux.Vars(r)
	amode := vars["mode"]

	alignment := scopeparser.NewAlignment(setMode(amode))

	ac, err := alignment.ParseMap()
	if err != nil {
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
	w.WriteHeader(http.StatusOK)
	w.Write(sendResponse(scopeResp))
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
