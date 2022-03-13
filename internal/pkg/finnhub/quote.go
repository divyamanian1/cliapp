package finnhub

import (
	"context"
	"log"

	fh "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

type Quote struct {
	// Open price of the day
	O *float32 `json:"o,omitempty"`
	// High price of the day
	H *float32 `json:"h,omitempty"`
	// Low price of the day
	L *float32 `json:"l,omitempty"`
	// Current price
	C *float32 `json:"c,omitempty"`
	// Previous close price
	Pc *float32 `json:"pc,omitempty"`
	// Change
	D *float32 `json:"d,omitempty"`
	// Percent change
	Dp *float32 `json:"dp,omitempty"`
}

func GetQuote(sym string, apikey string) (Quote, error) {
	var (
		q   fh.Quote
		err error
		q2  Quote
	)
	cfg := fh.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apikey)
	finnhubClient := fh.NewAPIClient(cfg).DefaultApi

	q, _, err = finnhubClient.Quote(context.Background()).Symbol(sym).Execute()
	if err != nil {
		log.Fatalln(err)
		q2.C = nil
		q2.Pc = nil
		q2.D = nil
		q2.Dp = nil
		q2.H = nil
		q2.L = nil
		q2.O = nil
	}
	q2.C = q.C
	q2.Pc = q.Pc
	q2.D = q.D
	q2.Dp = q.Dp
	q2.H = q.H
	q2.L = q.L
	q2.O = q.O
	return q2, err
}
