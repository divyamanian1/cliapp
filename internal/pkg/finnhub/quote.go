package finnhub

import (
	"context"

	fh "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

const apikey = "c8musp2ad3id1m4i6irg"

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

func GetQuote(sym string) (Quote, error) {
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
		q2 = Quote{
			C:  nil,
			Pc: nil,
			D:  nil,
			Dp: nil,
			H:  nil,
			L:  nil,
			O:  nil,
		}
	}
	q2 = Quote{
		C:  q.C,
		Pc: q.Pc,
		D:  q.D,
		Dp: q.Dp,
		H:  q.H,
		L:  q.L,
		O:  q.O,
	}

	return q2, err
}
