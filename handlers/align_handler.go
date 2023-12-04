package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ddefrancesco/scoperunner_server/etxclient"
	scopeparser "github.com/ddefrancesco/scoperunner_server/scopeparser"

	scopecommand "github.com/ddefrancesco/scoperunner_server/commands"
)

type scope_request struct {
	Mode string `json:"mode"`
}

type scope_response struct {
	Code     int    `json:"code"`
	Response string `json:"response"`
	Cmd      string `json:"cmd"`
}

type scope_err struct {
	Err            int    `json:"error_code"`
	ErrDescription string `json:"error_description"`
	ScopeFunction  string `json:"scope_function"`
	Cmd            string `json:"cmd"`
}

func sendResponse(sr etxclient.ScopeResponse) []byte {
	log.Println("sendResponse::Init -> eseguito")

	scoperesponse := scope_response{Code: http.StatusAccepted, Response: string(sr.Response), Cmd: sr.ExecCmd}

	jsonResponse, jsonError := json.Marshal(scoperesponse)

	if jsonError != nil {
		log.Println("Unable to encode JSON")
	}

	log.Println(string(jsonResponse))

	return jsonResponse
}

func AlignCommandHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("AlignCommandHandler::Init -> eseguito")
	// vars := mux.Vars(r)
	// amode := vars["mode"]
	var amode scope_request
	err := json.NewDecoder(r.Body).Decode(&amode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	alignment := scopeparser.NewAlignment(setMode(amode.Mode))

	ac, err := alignment.ParseMap()
	if err != nil {
		appErr := &scope_err{
			Err:            http.StatusBadRequest,
			ErrDescription: "Error parsing command: Opzione non valida",
			ScopeFunction:  "Align",
			Cmd:            amode.Mode,
		}
		JSONError(w, appErr, http.StatusBadRequest)
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

func JSONError(w http.ResponseWriter, err *scope_err, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
