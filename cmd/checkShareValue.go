/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	pnl "cliapp/internal/app/calculations"
	"cliapp/internal/app/quote"
	"cliapp/internal/pkg/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// checkShareValueCmd represents the checkShareValue command
var checkShareValueCmd = &cobra.Command{
	Use:   "checkShareValue",
	Short: "A brief description of your command",
	Long:  `Check Share Values and compare`,
	Run: func(cmd *cobra.Command, args []string) {
		checkShareValue()
	},
}

func init() {
	rootCmd.AddCommand(checkShareValueCmd)
}

func checkShareValue() {
	var (
		appcfg config.Config
		err    error
	)
	//hardcoded config file for now, flag with default ideal
	if appcfg, err = config.LoadFile("config/config.json"); err != nil {
		fmt.Println("unable to load config")
		os.Exit(1)
	}

	//Foreach share
	// check the value of C and PC and get the difference
	// multiply by 10 for 10 shares
	// output the value

	for _, s := range appcfg.Shares {
		q := quote.GetQuote(s, appcfg.Apikey)
		o := pnl.CalculatePL(*q.C, *q.Pc)
		fmt.Println(s, o)
	}
}