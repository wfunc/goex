package goex

import (
	"reflect"

	"github.com/wfunc/goex/v2/binance"
	"github.com/wfunc/goex/v2/httpcli"
	"github.com/wfunc/goex/v2/huobi"
	"github.com/wfunc/goex/v2/logger"
	"github.com/wfunc/goex/v2/okx"
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
