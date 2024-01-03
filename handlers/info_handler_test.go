package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestInfoCommandHandler(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/info/altitude", nil)
	req = mux.SetURLVars(req, map[string]string{"item": "altitude"})
	w := httptest.NewRecorder()
	InfoCommandHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	t.Logf("data: %v", string(data))
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if data == nil {
		t.Errorf("expected not null data  got %v", string(data))
	}
}
