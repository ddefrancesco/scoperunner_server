package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ddefrancesco/scoperunner_server/configurations"
	handler "github.com/ddefrancesco/scoperunner_server/handlers/commons"
	commons "github.com/ddefrancesco/scoperunner_server/models/commons"
	scopeparser "github.com/ddefrancesco/scoperunner_server/scopeparser"
)

func TestInitCommandHandler(t *testing.T) {

	t.Run("valid request", func(t *testing.T) {
		configurations.InitTestConfig()
		reqBody := `{"body": {"address":"Via Calcutta, Roma RM"}}`
		req := httptest.NewRequest(http.MethodPost, "/init", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		//InitCommandHandler(w, req)
		t.Log("InitCommandHandler::Init -> eseguito")

		var settings commons.ScopeInitRequest
		err := json.NewDecoder(req.Body).Decode(&settings.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		t.Logf("InitCommandHandler::Settings::Info -> %v ###", settings)

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
		res := w.Result()
		defer res.Body.Close()

		if strings.Contains(initCmd, ":SC") && strings.Contains(initCmd, ":SLs") && strings.Contains(initCmd, ":St") && strings.Contains(initCmd, ":Sg") {
			t.Log("InitCommandHandler::Command::Info -> OK")
		} else {
			t.Errorf("expected :SC<today>:#SLs<hour>:St<long>#:Sg0<lat>#, got %v", initCmd)
		}

	})

	t.Run("invalid request", func(t *testing.T) {
		t.Skip("TODO")
		configurations.InitTestConfig()
		reqBody := `{"body": "invalid"}`
		req := httptest.NewRequest(http.MethodPost, "/init", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		InitCommandHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status BadRequest, got %v", res.StatusCode)
		}

		var appErr commons.ScopeErr
		err := json.NewDecoder(res.Body).Decode(&appErr)
		if err != nil {
			t.Fatal(err)
		}

		if appErr.ScopeFunction != "Initialize" {
			t.Errorf("expected ScopeFunction Initialize, got %v", appErr.ScopeFunction)
		}

		if appErr.Err != http.StatusBadRequest {
			t.Errorf("expected Err BadRequest, got %v", appErr.Err)
		}
	})
}
