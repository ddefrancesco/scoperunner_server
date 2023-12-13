package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAlignCommandHandler(t *testing.T) {

	body := `{"body": "altaz"}`
	req := httptest.NewRequest(http.MethodPost, "/align", strings.NewReader(body))
	w := httptest.NewRecorder()
	AlignCommandHandler(w, req)
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
