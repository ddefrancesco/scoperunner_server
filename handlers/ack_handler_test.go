package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAckCommandHandler(t *testing.T) {

	req := httptest.NewRequest(http.MethodPost, "/ack", nil)
	w := httptest.NewRecorder()
	AckCommandHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if data == nil {
		t.Errorf("expected not null data  got %v", string(data))
	}
}
