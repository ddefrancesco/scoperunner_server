package scopeparser

import (
	"errors"
	"testing"
)

func TestNewInfoCommand(t *testing.T) {
	type args struct {
		m Info
	}
	tests := []struct {
		name string
		args args
		want *InfoCommand
	}{
		// TODO: Add test cases.
		{
			name: "test new info command",
			args: args{
				m: InfoAltitude,
			},
			want: &InfoCommand{
				Info: InfoAltitude,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInfoCommand(tt.args.m)
			if got.Info != tt.want.Info {
				t.Errorf("NewInfoCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfoCommand_InitMap(t *testing.T) {
	type fields struct {
		Info Info
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[Info]InfoCommandValue
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test initItems error",
			want: map[Info]InfoCommandValue{
				InfoAltitude: InfoAltitudeCmd,
				InfoAzimuth:  InfoAzimuthCmd,
				"accazzo":    "accazzo_cmd",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Skip()
			// ic := &InfoCommand{
			// 	Info: tt.fields.Info,
			// }
			// if got := ic.InitMap(); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("InfoCommand.InitMap() = %v, want %v", got, tt.want)

			// }
		})
	}
}

func TestInfoCommand_ParseMap(t *testing.T) {
	type fields struct {
		Info Info
	}
	tests := []struct {
		name    string
		fields  fields
		want    InfoCommandValue
		wantErr bool
	}{

		{
			name:   "test parse ok",
			fields: fields{Info: InfoAltitude},
			want:   InfoCommandValue(InfoAltitudeCmd),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := &InfoCommand{
				Info: tt.fields.Info,
			}
			got, err := ic.ParseMap()
			if (err != nil) != tt.wantErr {
				t.Errorf("InfoCommand.ParseMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InfoCommand.ParseMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfoCommand_StringValue(t *testing.T) {
	type fields struct {
		Info  Info
		Value InfoCommandValue
		Err   error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "test string value",
			fields: fields{Info: InfoAltitude, Value: InfoAltitudeCmd},
			want:   ":GA#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := &InfoCommand{
				Info:  tt.fields.Info,
				Value: tt.fields.Value,
				Err:   tt.fields.Err,
			}
			if got := ic.StringValue(); got != tt.want {
				t.Errorf("InfoCommand.StringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfoCommand_String(t *testing.T) {
	type fields struct {
		Info  Info
		Value InfoCommandValue
		Err   error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{
			name:   "test string",
			fields: fields{Info: InfoAltitude, Value: InfoAltitudeCmd},
			want:   "altitude",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := &InfoCommand{
				Info:  tt.fields.Info,
				Value: tt.fields.Value,
				Err:   tt.fields.Err,
			}
			if got := ic.String(); got != tt.want {
				t.Errorf("InfoCommand.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfoCommand_ErrUnknownCommand(t *testing.T) {
	ic := NewInfoCommand("accazzo")

	if _, err := ic.ParseMap(); !errors.Is(err, ErrUnknownInfoCommand) {
		t.Errorf("wrong error: %v", err)
	}

}
