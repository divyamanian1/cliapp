package quote

import (
	"reflect"
	"testing"
)

func TestGetQuote(t *testing.T) {
	type args struct {
		sym    string
		apikey string
	}
	tests := []struct {
		name string
		args args
		want Quote
	}{
		{
			name: "Incorrect API Key",
			args: args{
				sym:    "AAPL",
				apikey: "asd",
			},
			want: Quote{
				C:  nil,
				Pc: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQuote(tt.args.sym, tt.args.apikey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}
