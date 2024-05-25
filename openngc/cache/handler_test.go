package cache

import (
	"reflect"
	"testing"
)

func TestReadCsv(t *testing.T) {
	t.Skip()
	type args struct {
		csvFilePath string
	}
	tests := []struct {
		name string
		args args
		want NGCCatalog
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				csvFilePath: "../csv/test/valid.csv",
			},
			want: NGCCatalog{
				{
					Name:         "IC0001",
					Type:         "**",
					RA:           "00:08:27.05",
					Dec:          "+27:43:03.6",
					Const:        "Peg",
					MajAx:        0,
					MinAx:        0,
					PosAng:       0,
					BMag:         0,
					VMag:         0,
					JMag:         0,
					HMag:         0,
					KMag:         0,
					SurfBr:       0,
					Hubble:       "",
					Pax:          "",
					PmRA:         "",
					PmDec:        "",
					RadVel:       0,
					Redshift:     0,
					CstarUMag:    0,
					CstarBMag:    0,
					CstarVMag:    0,
					M:            "",
					NGC:          "",
					IC:           "",
					CstarNames:   "",
					Identifiers:  "",
					CommonNames:  "",
					NEDNotes:     "",
					OpenNGCNotes: "",
					Sources:      "Type:1|RA:1|Dec:1|Const:99",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadCsv(tt.args.csvFilePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCsv() = %v, want %v", got[0], tt.want[0])
			}
		})
	}
}
