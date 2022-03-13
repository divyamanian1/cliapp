package config

import (
	"reflect"
	"testing"
)

func TestLoadFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "Incorrect File",
			args: args{
				file: "xyz.json",
			},
			want: Config{
				Apikey: "",
				Shares: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
