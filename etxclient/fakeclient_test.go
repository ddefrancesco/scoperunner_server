package etxclient

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestFakeEtxClient_FetchQuery(t *testing.T) {
	type args struct {
		scopecmd string
	}
	tests := []struct {
		name string
		ec   *FakeEtxClient
		args args
		want interfaces.ETXResponse
	}{
		// TODO: Add test cases.
		{
			name: ":GC#",
			ec:   &FakeEtxClient{},
			args: args{
				scopecmd: ":GC#",
			},
			want: interfaces.ETXResponse{
				Err:      nil,
				Response: []byte(time.Now().Format("01/02/06#")),
				ExecCmd:  ":GC#",
			},
		},
		//Scope: "scope -f /home/ddefrancesco/go/src/github.com/ddefrancesco/scoperunner_server/etxclient/test_data/test_scope.txt",

		//scopecmd: "scope -f /home/ddefrancesco/go/src/github.com/ddefrancesco/scoperunner_server/etxclient/test_data/test_scope.txt",

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec := &FakeEtxClient{}
			if got := ec.FetchQuery(tt.args.scopecmd); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeEtxClient.FetchQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFakeEtxClient_Date_Format(t *testing.T) {
	layout := "01/02/06#"
	//date := time.Now().Format(layout)
	date := time.Date(2024, 1, 14, 14, 30, 45, 100, time.Local).Format(layout)
	assert.Equal(t, date, "01/14/24#")
}

func TestFakeEtxClient_MinQuality(t *testing.T) {
	t.Skip("Not meaningful")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	assert.True(t, r1.Intn(7) >= 1 && r1.Intn(7) <= 7)

}
