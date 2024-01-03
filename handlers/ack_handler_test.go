package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAckCommandHandler0(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/ack", nil)
	w := httptest.NewRecorder()
	AckCommandHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if data == nil {
		t.Errorf("expected not null data  got %v", string(data))
	}
}

func TestAckCommandHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test ok",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/ack", nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AckCommandHandler(tt.args.w, tt.args.r)
		})
	}
}
