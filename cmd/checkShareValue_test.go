/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import "testing"

func BenchmarkCheckShareValueSingle(b *testing.B) {
	args := []string{"AAPL"}
	checkShareValue(args)
}

func BenchmarkCheckShareValueTwo(b *testing.B) {
	args := []string{"AAPL", "MSFT"}
	checkShareValue(args)
}

func Test_checkShareValue(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Successful",
			args: args{
				[]string{"AAPL"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checkShareValue(tt.args.args)
		})
	}

}
