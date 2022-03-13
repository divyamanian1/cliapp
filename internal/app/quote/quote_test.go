package quote

import (
	"reflect"
	"testing"
)

func TestGetQuote(t *testing.T) {
	type args struct {
		sym string
	}
	tests := []struct {
		name string
		args args
		want Quote
	}{
		{
			name: "Unsuccessful",
			args: args{
				sym: "9",
			},
			want: Quote{
				C:  new(float32),
				Pc: new(float32),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQuote(tt.args.sym); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}
