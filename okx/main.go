package okx

import (
	"github.com/wfunc/goex/v2/okx/futures"
	"github.com/wfunc/goex/v2/okx/spot"
)

type OKx struct {
	Spot    *spot.Spot
	Futures *futures.Futures
	Swap    *futures.Swap
}

func New() *OKx {
	return &OKx{
		Spot:    spot.New(),
		Futures: futures.New(),
		Swap:    futures.NewSwap(),
	}
}
