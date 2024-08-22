package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	configuration "github.com/ddefrancesco/scoperunner_server/configurations"
)

func TestSetCommandHandler(t *testing.T) {
	err := configuration.InitTestConfig()
	if err != nil {
		t.Fatalf("Failed to initialize test config: %v", err)
	}

	tests := []struct {
		name           string
		body           string
		expectedStatus int
	}{
		{
			name:           "Valid set command",
			body:           `{"body": {"target_dec": "+46*01"}}`,
			expectedStatus: http.StatusAccepted,
		},
		{
			name:           "Invalid JSON",
			body:           `{"body": {"command": "set", "parameter": "time", "value": "20:30:00"`,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Missing parameter",
			body:           `{"body": {"" : "20:30:00"}}`,
			expectedStatus: http.StatusBadRequest,
		},
		// {
		// 	name:           "Missing value",
		// 	body:           `{"body": {"target_dec": ""}}`,
		// 	expectedStatus: http.StatusBadRequest,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/set", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			SetCommandHandler(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			data, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}

			if tt.expectedStatus == http.StatusOK {
				var response map[string]interface{}
				err = json.Unmarshal(data, &response)
				if err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if response["status"] != "success" {
					t.Errorf("expected status 'success', got %v", response["status"])
				}
			}
		})
	}
}
