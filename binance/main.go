package binance

import (
	"github.com/wfunc/goex/v2/binance/futures/fapi"
	"github.com/wfunc/goex/v2/binance/spot"
)

type Binance struct {
	Spot *spot.Spot
	Swap *fapi.FApi
}

func New() *Binance {
	return &Binance{
		Spot: spot.New(),
		Swap: fapi.NewFApi(),
	}
}
