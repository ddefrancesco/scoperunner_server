package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	configuration "github.com/ddefrancesco/scoperunner_server/configurations"
	scopeparser "github.com/ddefrancesco/scoperunner_server/scopeparser"
)

func TestAlignCommandHandler(t *testing.T) {
	err := configuration.InitTestConfig()
	if err != nil {
		panic(err)
	}
	body := `{"body": {"mode": "altaz"}}`
	req := httptest.NewRequest(http.MethodPost, "/align", strings.NewReader(body))
	w := httptest.NewRecorder()
	AlignCommandHandler(w, req)
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

func Test_setMode(t *testing.T) {
	type args struct {
		mode string
	}
	tests := []struct {
		name string
		args args
		want scopeparser.AlignMode
	}{

		{
			name: "Case AZ",
			args: args{mode: "altaz"},
			want: scopeparser.AltAz,
		},
		{
			name: "Case Polar",
			args: args{mode: "polar"},
			want: scopeparser.Polar,
		},
		{
			name: "Case Land",
			args: args{mode: "land"},
			want: scopeparser.Land,
		},
		{
			name: "Case Error",
			args: args{mode: "pipppo"},
			want: "error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setMode(tt.args.mode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
