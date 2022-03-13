/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	pnl "cliapp/internal/app/calculations"
	"cliapp/internal/app/quote"
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

// checkShareValueCmd represents the checkShareValue command
var checkShareValueCmd = &cobra.Command{
	Use:   "checkShareValue",
	Short: "A brief description of your command",
	Long:  `Check Share Values and compare`,
	Run: func(cmd *cobra.Command, args []string) {
		checkShareValue(args)
	},
}

func init() {
	rootCmd.AddCommand(checkShareValueCmd)
}

func checkShareValue(args []string) {

	//Foreach share
	// check the value of C and PC and get the difference
	// multiply by 10 for 10 shares
	// output the value

	var wg sync.WaitGroup
	if len(args) == 0 {
		args = []string{"MSFT", "AAPL"}
	}

	for _, s := range args {
		wg.Add(1)
		s2 := s
		go func() {
			defer wg.Done()
			q := quote.GetQuote(s2)
			o := pnl.CalculatePL(*q.C, *q.Pc)
			fmt.Println(s2, o)
		}()
	}
	wg.Wait()
}
