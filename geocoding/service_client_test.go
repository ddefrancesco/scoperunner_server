package geocoding

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	cache "github.com/ddefrancesco/scoperunner_server/cache"
	configuration "github.com/ddefrancesco/scoperunner_server/configurations"
)

func TestGetAutostarLocation(t *testing.T) {
	err := configuration.InitConfig()
	if err != nil {
		panic(err)
	}
	gcache := cache.New[string, *AutostarLatLong]()
	type args struct {
		address  Address
		geoCache *cache.Cache[string, *AutostarLatLong]
	}
	tests := []struct {
		name    string
		args    args
		want    *AutostarLatLong
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test ok",
			args: args{
				address:  Address{Location: "Via Calcutta, Roma RM"},
				geoCache: gcache,
			},
			want:    &AutostarLatLong{AutostarLat: "41*49", AutostarLong: "012*26"},
			wantErr: false,
		},

		{
			name: "test cache ok",
			args: args{
				address:  Address{Location: "Via Calcutta, Roma RM"},
				geoCache: gcache,
			},
			want:    &AutostarLatLong{AutostarLat: "41*49", AutostarLong: "012*26"},
			wantErr: false,
		},

		// {
		// 	name: "test ko",
		// 	args: args{
		// 		address:  Address{Location: "a"},
		// 		geoCache: gcache,
		// 	},
		// 	want:    &AutostarLatLong{AutostarLat: "91*49", AutostarLong: "181*26"},
		// 	wantErr: true,
		// },
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			got, err := GetAutostarLocation(tt.args.address, tt.args.geoCache)
			end := time.Now()
			t.Logf("Test "+strconv.Itoa(i)+" took %f seconds.\n", end.Sub(start).Seconds())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAutostarLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAutostarLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
