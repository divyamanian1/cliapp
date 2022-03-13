package quote

import (
	"cliapp/internal/pkg/finnhub"
)

//TBD Split this into getting just current or just previous closing price
// in order to truly remove dependency from finnhub

type Quote struct {
	// Current price
	C *float32 `json:"c,omitempty"`
	// Previous close price
	Pc *float32 `json:"pc,omitempty"`
}

func GetQuote(sym string, apikey string) Quote {
	var (
		q    finnhub.Quote
		e    error
		appq Quote
	)
	q, e = finnhub.GetQuote(sym, apikey)
	if e != nil {
		appq.C = nil
		appq.Pc = nil
	}
	appq.C = q.C
	appq.Pc = q.Pc
	return appq
}
