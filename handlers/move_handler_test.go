package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ddefrancesco/scoperunner_server/configurations"
	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	"github.com/ddefrancesco/scoperunner_server/models/commons"
)

func TestMoveCommandHandler(t *testing.T) {
	configurations.InitTestConfig()

	t.Run("valid request slew_north", func(t *testing.T) {
		reqBody := `{"body": {"slew_north": ""}}`
		req := httptest.NewRequest(http.MethodPost, "/move", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		GotoCommandHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusAccepted {
			t.Errorf("expected status Accepted, got %v", res.StatusCode)
		}
	})

	t.Run("valid request slew_south", func(t *testing.T) {
		reqBody := `{"body": {"slew_south": ""}}`
		req := httptest.NewRequest(http.MethodPost, "/move", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		GotoCommandHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusAccepted {
			t.Errorf("expected status Accepted, got %v", res.StatusCode)
		}
	})

	t.Run("valid request slew_east", func(t *testing.T) {
		reqBody := `{"body": {"slew_east": ""}}`
		req := httptest.NewRequest(http.MethodPost, "/move", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		GotoCommandHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusAccepted {
			t.Errorf("expected status Accepted, got %v", res.StatusCode)
		}
	})

	t.Run("valid request slew_west", func(t *testing.T) {
		reqBody := `{"body": {"slew_west": ""}}`
		req := httptest.NewRequest(http.MethodPost, "/move", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		GotoCommandHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusAccepted {
			t.Errorf("expected status Accepted, got %v", res.StatusCode)
		}
	})
	t.Run("valid request slew@ rad dec", func(t *testing.T) {
		reqBody := `{"body": {"slew_at_radec": ""}}`
		req := httptest.NewRequest(http.MethodPost, "/move", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		GotoCommandHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusAccepted {
			t.Errorf("expected status Accepted, got %v", res.StatusCode)
		}
	})

	t.Run("valid request slew@ target", func(t *testing.T) {
		reqBody := `{"body": {"slew_at_target": ""}}`
		req := httptest.NewRequest(http.MethodPost, "/move", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		GotoCommandHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusAccepted {
			t.Errorf("expected status Accepted, got %v", res.StatusCode)
		}
	})

	t.Run("valid request slew_west for 1 sec", func(t *testing.T) {
		reqBody := `{"body": {"slew_west_ms": "1000"}}`
		req := httptest.NewRequest(http.MethodPost, "/move", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		GotoCommandHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusAccepted {
			t.Errorf("expected status Accepted, got %v", res.StatusCode)
		}
	})

	t.Run("invalid request", func(t *testing.T) {
		reqBody := `{"body": {"slew_neast": ""}}`
		req := httptest.NewRequest(http.MethodPost, "/move", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		GotoCommandHandler(w, req)

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

		if appErr.ScopeFunction != "Move" {
			t.Errorf("expected ScopeFunction Move, got %v", appErr.ScopeFunction)
		}

		if appErr.Err != http.StatusBadRequest {
			t.Errorf("expected Err BadRequest, got %v", appErr.Err)
		}
	})

	t.Run("serial device error", func(t *testing.T) {
		t.Skip()
		reqBody := `{"body": {"RM": "123", "RG": "456"}}`
		req := httptest.NewRequest(http.MethodPost, "/move", bytes.NewBufferString(reqBody))
		w := httptest.NewRecorder()

		// Mock the serial device to return an error
		// serialDeviceMock := &serialDeviceMock{
		// 	ExecCommandFunc: func(command string) interfaces.ETXResponse {
		// 		return interfaces.ETXResponse{
		// 			Err:      &serial.PortError{},
		// 			Response: []byte("Command Failed"),
		// 			ExecCmd:  "move",
		// 		}
		// 	},
		// }
		//handler.SetScopeClient(serialDeviceMock)

		GotoCommandHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != http.StatusInternalServerError {
			t.Errorf("expected status InternalServerError, got %v", res.StatusCode)
		}
	})
}

type serialDeviceMock struct {
	ExecCommandFunc func(command string) interfaces.ETXResponse
}

func (m *serialDeviceMock) ExecCommand(command string) interfaces.ETXResponse {
	return m.ExecCommandFunc(command)
}
