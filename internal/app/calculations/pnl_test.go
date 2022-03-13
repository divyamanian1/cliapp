package pnl

import "testing"

func TestCalculatePL(t *testing.T) {
	type args struct {
		c  float32
		pc float32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Profit",
			args: args{
				c:  100,
				pc: 80,
			},
			want: "Profit made for 10 shares: 200.00.",
		},
		{
			name: "Loss",
			args: args{
				c:  80,
				pc: 100,
			},
			want: "Loss made for 10 shares: 200.00.",
		},
		{
			name: "No profit no loss",
			args: args{
				c:  100,
				pc: 100,
			},
			want: "No profit or loss",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatePL(tt.args.c, tt.args.pc); got != tt.want {
				t.Errorf("CalculatePL() = %v, want %v", got, tt.want)
			}
		})
	}
}
