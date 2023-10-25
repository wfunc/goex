package goex

import (
	"reflect"

	"github.com/wfunc/goex/binance"
	"github.com/wfunc/goex/httpcli"
	"github.com/wfunc/goex/huobi"
	"github.com/wfunc/goex/logger"
	"github.com/wfunc/goex/okx"
)

var (
	DefaultHttpCli = httpcli.Cli
)

var (
	OKx     = okx.New()
	Binance = binance.New()
	HuoBi   = huobi.New()
)

func SetDefaultHttpCli(cli httpcli.IHttpClient) {
	logger.Infof("use new http client implement: %s", reflect.TypeOf(cli).Elem().String())
	httpcli.Cli = cli
}
