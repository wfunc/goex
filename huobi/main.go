package huobi

import (
	"github.com/wfunc/goex/huobi/futures"
	"github.com/wfunc/goex/huobi/spot"
)

type HuoBi struct {
	Spot    *spot.Spot
	Futures *futures.Futures
}

func New() *HuoBi {
	return &HuoBi{
		Spot:    spot.New(),
		Futures: futures.New(),
	}
}
