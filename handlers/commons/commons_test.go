package commons

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	"github.com/ddefrancesco/scoperunner_server/models/commons"
	"github.com/spf13/viper"
)

func TestSendResponse(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	response := interfaces.ETXResponse{
		Response: []byte("test response"),
		ExecCmd:  "test command",
	}

	result := SendResponse(req, response)

	var scopeResponse commons.ScopeResponse
	err := json.Unmarshal(result, &scopeResponse)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if scopeResponse.Code != http.StatusOK {
		t.Errorf("expected code %d, got %d", http.StatusOK, scopeResponse.Code)
	}
	if scopeResponse.Response != "test response" {
		t.Errorf("expected response 'test response', got '%s'", scopeResponse.Response)
	}
	if scopeResponse.Cmd != "test command" {
		t.Errorf("expected cmd 'test command', got '%s'", scopeResponse.Cmd)
	}
}

func TestSendResponsePost(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/test", nil)
	response := interfaces.ETXResponse{
		Response: []byte("post response"),
		ExecCmd:  "post command",
	}

	result := SendResponse(req, response)

	var scopeResponse commons.ScopeResponse
	err := json.Unmarshal(result, &scopeResponse)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if scopeResponse.Code != http.StatusAccepted {
		t.Errorf("expected code %d, got %d", http.StatusAccepted, scopeResponse.Code)
	}
}

func TestSendResponses(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	responses := []interfaces.ETXResponse{
		{Response: []byte("response1"), ExecCmd: "cmd1"},
		{Response: []byte("response2"), ExecCmd: "cmd2"},
	}

	result := SendResponses(req, responses)

	var scopeResponses []commons.ScopeResponse
	err := json.Unmarshal(result, &scopeResponses)
	if err != nil {
		t.Fatalf("failed to unmarshal responses: %v", err)
	}

	if len(scopeResponses) != 2 {
		t.Errorf("expected 2 responses, got %d", len(scopeResponses))
	}
	if scopeResponses[0].Response != "response1" {
		t.Errorf("expected first response 'response1', got '%s'", scopeResponses[0].Response)
	}
	if scopeResponses[1].Cmd != "cmd2" {
		t.Errorf("expected second cmd 'cmd2', got '%s'", scopeResponses[1].Cmd)
	}
}

func TestJSONError(t *testing.T) {
	w := httptest.NewRecorder()
	scopeErr := &commons.ScopeErr{
		Err:            500,
		ErrDescription: "test error",
		ScopeFunction:  "test function",
		Cmd:            "test cmd",
	}

	JSONError(w, scopeErr, http.StatusInternalServerError)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("expected JSON content type, got %s", contentType)
	}

	var result commons.ScopeErr
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("failed to unmarshal error response: %v", err)
	}

	if result.Err != 500 {
		t.Errorf("expected error code 500, got %d", result.Err)
	}
	if result.ErrDescription != "test error" {
		t.Errorf("expected error description 'test error', got '%s'", result.ErrDescription)
	}
}

func TestGetScopeClientFake(t *testing.T) {
	viper.Set("environments.fakescope", true)
	defer viper.Reset()

	client := GetScopeClient()
	if client == nil {
		t.Fatal("GetScopeClient returned nil")
	}
}

func TestGetScopeClientReal(t *testing.T) {
	viper.Set("environments.fakescope", false)
	defer viper.Reset()

	client := GetScopeClient()
	if client == nil {
		t.Fatal("GetScopeClient returned nil")
	}
}