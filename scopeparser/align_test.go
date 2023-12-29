package scopeparser

import (
	"reflect"
	"testing"
)

func TestNewAlignment(t *testing.T) {
	type args struct {
		m AlignMode
	}
	tests := []struct {
		name string
		args args
		want *Alignment
	}{

		{
			name: "test NewAlignment AZ",
			args: args{m: AltAz},
			want: &Alignment{mode: AltAz},
		},
		{
			name: "test NewAlignment Polar",
			args: args{m: Polar},
			want: &Alignment{mode: Polar},
		},
		{
			name: "test NewAlignment Land",
			args: args{m: Land},
			want: &Alignment{mode: Land},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlignment(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlignment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initItems(t *testing.T) {
	tests := []struct {
		name string
		want map[AlignMode]AlignCommandValue
	}{

		{
			name: "test initItems ok",
			want: map[AlignMode]AlignCommandValue{
				AltAz: AltAzCmd,
				Polar: PolarCmd,
				Land:  LandCmd,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initItems(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlignment_ParseMap(t *testing.T) {
	type fields struct {
		mode AlignMode
	}
	tests := []struct {
		name    string
		fields  fields
		want    AlignCommandValue
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "test ParseMap ok",
			fields: fields{mode: AltAz},
			want:   AltAzCmd,
		},
		{
			name:    "test ParseMap error",
			fields:  fields{mode: "asdf"},
			want:    "error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Alignment{
				mode: tt.fields.mode,
			}
			got, err := p.ParseMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("Alignment.ParseMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Alignment.ParseMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
