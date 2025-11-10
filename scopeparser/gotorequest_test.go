package scopeparser

import (
	"reflect"
	"strings"
	"testing"

	"github.com/ddefrancesco/scoperunner_server/configurations"
)

func TestNewGotoRequest(t *testing.T) {

	m := map[string]string{"target_ra": "12:34:56", "target_dec": "+45:67:89"}
	request := NewGotoRequest(m)
	if !reflect.DeepEqual(request.Goto, m) {
		t.Errorf("NewGotoRequest() = %v, want %v", request.Goto, m)
	}
}

func TestGotoRequest_ParseMap(t *testing.T) {
	request := &GotoRequest{
		Goto: map[string]string{"target_ra": "12:34:56", "target_dec": "+45:67:89"},
	}
	expected := map[string]string{"Sr": "12:34:56", "Sd": "+45:67:89"}
	result, err := request.ParseMap()
	if err != nil {
		t.Errorf("ParseMap() error = %v, wantErr %v", err, nil)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseMap() = %v, want %v", result, expected)
	}
}

func TestGotoRequest_CheckGotoRADecCommand(t *testing.T) {
	request := &GotoRequest{}
	expected := ":MS#"
	cmd, err := request.CheckGotoRADecCommand()
	if err != nil {
		t.Errorf("CheckGotoRADecCommand() error = %v, wantErr %v", err, nil)
	}
	if cmd != expected {
		t.Errorf("CheckGotoRADecCommand() = %v, want %v", cmd, expected)
	}
}
func TestGotoRequest_FindDeepSpaceObjectCommand(t *testing.T) {
	configurations.InitTestConfig()
	request := &GotoRequest{
		Goto: map[string]string{"goto": "NGC0224"},
	}
	cmd, err := request.FindDeepSpaceObjectCommand()
	if err != nil {
		t.Errorf("FindDeepSpaceObjectCommand() error = %v, wantErr %v", err, nil)
	}
	if !strings.HasPrefix(cmd, ":Sr") || !strings.Contains(cmd, "#:Sd") {
		t.Errorf("FindDeepSpaceObjectCommand() = %v, want command with :Sr and :Sd", cmd)
	}
}
