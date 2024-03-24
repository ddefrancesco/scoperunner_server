package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ddefrancesco/scoperunner_server/etxclient"
	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	"github.com/ddefrancesco/scoperunner_server/models/commons"
	"github.com/spf13/viper"
)

func SendResponse(r *http.Request, sr interfaces.ETXResponse) []byte {
	log.Println("sendResponse::Init -> eseguito")
	status := http.StatusOK
	if r.Method == http.MethodPost {
		status = http.StatusAccepted
	}
	scoperesponse := commons.ScopeResponse{Code: status, Response: string(sr.Response), Cmd: sr.ExecCmd}

	jsonResponse, jsonError := json.Marshal(scoperesponse)

	if jsonError != nil {
		log.Println("Unable to encode JSON")
	}

	log.Println(string(jsonResponse))

	return jsonResponse
}

func SendResponses(r *http.Request, sr []interfaces.ETXResponse) []byte {
	log.Println("sendResponse::Init -> eseguito")
	status := http.StatusOK
	var jsonResponses []commons.ScopeResponse
	for _, v := range sr {
		scoperesponse := commons.ScopeResponse{Code: status, Response: string(v.Response), Cmd: v.ExecCmd}
		jsonResponses = append(jsonResponses, scoperesponse)
	}
	jsonResponse, jsonError := json.Marshal(jsonResponses)

	if jsonError != nil {
		log.Println("Unable to encode JSON")
	}

	return jsonResponse
}

func JSONError(w http.ResponseWriter, err *commons.ScopeErr, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

func GetScopeClient() interfaces.SerialClient {
	var serialDevice interfaces.SerialClient
	if viper.GetBool("environments.fakescope") {
		etx := etxclient.NewFakeClient()
		serialDevice = etx
	} else {
		etx := etxclient.NewClient()
		serialDevice = etx
	}
	log.Println("GetScopeClient::Init -> eseguito")
	return serialDevice
}
