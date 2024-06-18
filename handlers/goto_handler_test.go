package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ddefrancesco/scoperunner_server/configurations"
	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	"github.com/ddefrancesco/scoperunner_server/scopeparser"
)

func TestGotoCommandHandler202(t *testing.T) {
	configurations.InitTestConfig()
	// Test case: Valid goto request
	body := `{"body": {"goto": "NGC0224"}}`
	req := httptest.NewRequest(http.MethodPost, "/goto", bytes.NewBufferString(body))
	w := httptest.NewRecorder()

	GotoCommandHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Expected status code %d, got %d", http.StatusAccepted, resp.StatusCode)
	}

}

func TestGotoCommandHandler400(t *testing.T) {
	t.Skip("Per ora lasciamo stare")

	// Test case: Invalid goto request
	invalidBody := `{"body": {"goto": "invalid"}}`
	req := httptest.NewRequest(http.MethodPost, "/goto", bytes.NewBufferString(invalidBody))
	w := httptest.NewRecorder()

	GotoCommandHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

// func TestGotoCommandHandlerWithMocks(t *testing.T) {
// 	// Mock the GetScopeClient function
// 	originalGetScopeClient := commons
// 	defer func() { GetScopeClient = originalGetScopeClient }()

// 	mockClient := &MockScopeClient{}
// 	GetScopeClient = func() interfaces.ScopeClient {
// 		return mockClient
// 	}

// 	// Test case: Valid goto request with mock client
// 	body := `{"body": {"goto": "M31"}}`
// 	req := httptest.NewRequest(http.MethodPost, "/goto", bytes.NewBufferString(body))
// 	w := httptest.NewRecorder()

// 	GotoCommandHandler(w, req)

// 	resp := w.Result()
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusAccepted {
// 		t.Errorf("Expected status code %d, got %d", http.StatusAccepted, resp.StatusCode)
// 	}

// 	if !mockClient.ExecCommandCalled {
// 		t.Error("ExecCommand was not called on the mock client")
// 	}
// }

type MockScopeClient struct {
	ExecCommandCalled bool
}

func (m *MockScopeClient) ExecCommand(command string) interfaces.ETXResponse {
	m.ExecCommandCalled = true
	return interfaces.ETXResponse{}
}

func TestNewGotoRequest(t *testing.T) {
	// Test case: Valid goto request
	body := scopeparser.GotoRequest{
		Goto: map[string]string{
			"goto": "M31",
		},
	}
	gotoRequest := scopeparser.NewGotoRequest(body.Goto)
	if gotoRequest == nil {
		t.Error("Expected non-nil GotoRequest")
	}

	// Test case: Invalid goto request
	invalidBody := scopeparser.GotoRequest{
		Goto: map[string]string{
			"invalid": "value",
		},
	}
	gotoRequest = scopeparser.NewGotoRequest(invalidBody.Goto)
	if gotoRequest != nil {
		t.Error("Expected nil GotoRequest for invalid request")
	}
}
