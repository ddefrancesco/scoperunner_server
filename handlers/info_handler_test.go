package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	converter "github.com/ddefrancesco/scoperunner_server/commands/converters"
	configuration "github.com/ddefrancesco/scoperunner_server/configurations"
	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	handler "github.com/ddefrancesco/scoperunner_server/handlers/commons"
	"github.com/ddefrancesco/scoperunner_server/scopeparser"
)

func TestInfoCommandHandler(t *testing.T) {
	t.Skip(t.Name())
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

func TestInfoCommandHandlerParseRequestParams(t *testing.T) {
	err := configuration.InitConfig()
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodGet, "/info", nil)
	req = mux.SetURLVars(req, map[string]string{"infos": "altitude,azimuth,declination"})
	vars := mux.Vars(req)
	var infoParam = vars["infos"]
	t.Logf("params: %v", string(infoParam))

	tMap, _ := converter.RequestParamsToInfoArray(infoParam)
	var infoCommandArray []scopeparser.InfoCommand
	var scopeResponses []interfaces.ETXResponse
	for _, v := range tMap {
		t.Logf("value: %v", v)
		if v == scopeparser.InfoAzimuth || v == scopeparser.InfoAltitude || v == scopeparser.InfoDeclination {
			infoCommand := *scopeparser.NewInfoCommand(v)
			infoCommandArray = append(infoCommandArray, infoCommand)

		}

	}
	assert.Equal(t, 3, len(infoCommandArray))
	serialDevice := handler.GetScopeClient()
	for _, v := range infoCommandArray {
		t.Logf("infoCommnd: %v", v)
		ic, _ := v.ParseMap()
		t.Logf("infoCommndValue: %v", ic)
		scopeResponse := serialDevice.ExecCommand(v.StringValue())
		scopeResponses = append(scopeResponses, scopeResponse)
	}
	assert.Equal(t, 3, len(scopeResponses))
	assert.Contains(t, scopeResponses, interfaces.ETXResponse{
		Err:      nil,
		Response: []byte("s41*53’30#"),
		ExecCmd:  ":GA#",
	})

}
