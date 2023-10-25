### Introduction
* Unified and standardize the interfaces of various cryptocurrency trading platforms.

![goex](goex_struct.png)

### Donate
* [BTC] 1GoEXwVvXG7kNdQSFaUNF35A3izHojLGxP
* [USDT-TRC20] TGoExC6xvzE4wSA9cYZnwcPaXEjibA5Vtc

### Example

```golang
package main

import (
	goexv2 "github.com/wfunc/goex"
	"github.com/wfunc/goex/logger"
	"github.com/wfunc/goex/model"
	"github.com/wfunc/goex/options"
	"log"
)

func main() {
	logger.SetLevel(logger.DEBUG)                             //设置日志输出级别
	//goexv2.DefaultHttpCli.SetProxy("socks5://127.0.0.1:1080") //socks代理
	goexv2.DefaultHttpCli.SetTimeout(5)                       // 5 second

	_, _, err := goexv2.OKx.Spot.GetExchangeInfo() //建议调用
	if err != nil {
		panic(err)
	}
	btcUSDTCurrencyPair, err := goexv2.OKx.Spot.NewCurrencyPair(model.BTC, model.USDT)//建议这样构建CurrencyPair
	if err != nil {
		panic(err)
	}
		
	//共有api调用
	log.Println(goexv2.OKx.Spot.GetTicker(btcUSDTCurrencyPair))

	//私有API调用
	okxPrvApi := goexv2.OKx.Spot.NewPrvApi(
		options.WithApiKey(""), 
		options.WithApiSecretKey(""), 
		options.WithPassphrase(""))
	
	//创建订单
	order, _, err := okxPrvApi.CreateOrder(btcUSDTCurrencyPair, 0.01, 18000, model.Spot_Buy, model.OrderType_Limit)
	log.Println(err)
	log.Println(order)
}
```

### Thanks
<a href="https://www.jetbrains.com/?from=goex"><img src="https://account.jetbrains.com/static/images/jetbrains-logo-inv.svg" height="120" alt="JetBrains"/></a>

