package converters

import (
	"reflect"
	"testing"

	"github.com/ddefrancesco/scoperunner_server/scopeparser"
)

// func TestRequestParamsToInfoArray(t *testing.T) {
// 	type args struct {
// 		params string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []scopeparser.Info
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "test",
// 			args: args{
// 				params: "altitude,azimuth,browse_bml",
// 				// params: "a=b&c=d&e=f&g=h",
// 				// params: "a=b&c=d&e=f&g=h&i=j&k=l&m=n&o=p&q=r&s=t&u=v&w=x&y=z",
// 			},
// 			want: []scopeparser.Info{scopeparser.InfoAltitude, scopeparser.InfoAzimuth, scopeparser.InfoBrighterMagLimit},nil,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := RequestParamsToInfoArray(tt.args.params); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("RequestParamsToInfoArray() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestRequestParamsToInfoArray(t *testing.T) {
	type args struct {
		params string
	}
	tests := []struct {
		name    string
		args    args
		want    []scopeparser.Info
		wantErr bool
	}{

		{
			name: "test ok",
			args: args{
				params: "altitude,azimuth,browse_bml",
				// params: "a=b&c=d&e=f&g=h",
				// params: "a=b&c=d&e=f&g=h&i=j&k=l&m=n&o=p&q=r&s=t&u=v&w=x&y=z",
			},
			want:    []scopeparser.Info{scopeparser.InfoAltitude, scopeparser.InfoAzimuth, scopeparser.InfoBrighterMagLimit},
			wantErr: false,
		},
		{
			name: "test ko",
			args: args{
				params: "altitude,azimuth,accazzo",
				// params: "a=b&c=d&e=f&g=h",
				// params: "a=b&c=d&e=f&g=h&i=j&k=l&m=n&o=p&q=r&s=t&u=v&w=x&y=z",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RequestParamsToInfoArray(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestParamsToInfoArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestParamsToInfoArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
