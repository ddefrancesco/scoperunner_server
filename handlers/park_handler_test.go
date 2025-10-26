package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ddefrancesco/scoperunner_server/configurations"
	commons "github.com/ddefrancesco/scoperunner_server/models/commons"
)

func TestParkCommandHandler(t *testing.T) {
	configurations.InitTestConfig()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		//setup   func()
	}{
		{
			name: "Valid parking request",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/park", strings.NewReader(`{"body":{"parking":"seek"}}`)),
			},
			wantErr: false,
			// setup: func() {
			// 	handler.SetScopeClient(&mockScopeClient{})
			// },
		},
		{
			name: "Invalid JSON body",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/park", strings.NewReader(`{"body":{"parking":}`)),
			},
			wantErr: true,
		},
		{
			name: "Invalid parking option",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/park", strings.NewReader(`{"body":{"parking":"invalid"}}`)),
			},
			wantErr: true,
		},
		{
			name: "Serial device error",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/park", strings.NewReader(`{"body":{"parking":"home"}}`)),
			},
			wantErr: true,
			// setup: func() {
			// 	handler.SetScopeClient(&mockScopeClientError{})
			// },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if tt.setup != nil {
			// 	tt.setup()
			// }

			ParkCommandHandler(tt.args.w, tt.args.r)

			resp := tt.args.w.(*httptest.ResponseRecorder)
			if tt.wantErr {
				if resp.Code == http.StatusAccepted {
					t.Errorf("ParkCommandHandler() expected error, got status %v", resp.Code)
				}
			} else {
				if resp.Code != http.StatusAccepted {
					t.Errorf("ParkCommandHandler() got status %v, want %v", resp.Code, http.StatusAccepted)
				}
			}
		})
	}
}

type mockScopeClient struct{}

func (m *mockScopeClient) ExecCommand(cmd string) commons.ScopeResponse {
	return commons.ScopeResponse{Response: "OK"}
}

type mockScopeClientError struct{}

// func (m *mockScopeClientError) ExecCommand(cmd string) commons.ScopeResponse {
// 	return commons.ScopeResponse{Err: errors.New("serial device error")}
// }
