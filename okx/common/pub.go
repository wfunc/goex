package common

import (
	"fmt"
	"net/http"
	"net/url"

	. "github.com/wfunc/goex/httpcli"
	"github.com/wfunc/goex/logger"
	. "github.com/wfunc/goex/model"
	. "github.com/wfunc/goex/util"
)

func (okx *OKxV5) GetName() string {
	return "okx.com"
}

func (okx *OKxV5) GetDepth(pair CurrencyPair, size int, opt ...OptionParameter) (*Depth, []byte, error) {
	params := url.Values{}
	params.Set("instId", pair.Symbol)
	params.Set("sz", fmt.Sprint(size))
	MergeOptionParams(&params, opt...)

	data, responseBody, err := okx.DoNoAuthRequest("GET", okx.UriOpts.Endpoint+okx.UriOpts.DepthUri, &params)
	if err != nil {
		return nil, responseBody, err
	}

	dep, err := okx.UnmarshalOpts.DepthUnmarshaler(data)
	if err != nil {
		return nil, data, err
	}

	dep.Pair = pair

	return dep, responseBody, err
}

func (okx *OKxV5) GetTicker(pair CurrencyPair, opt ...OptionParameter) (*Ticker, []byte, error) {
	params := url.Values{}
	params.Set("instId", pair.Symbol)

	data, responseBody, err := okx.DoNoAuthRequest("GET", okx.UriOpts.Endpoint+okx.UriOpts.TickerUri, &params)
	if err != nil {
		return nil, data, err
	}

	tk, err := okx.UnmarshalOpts.TickerUnmarshaler(data)
	if err != nil {
		return nil, nil, err
	}

	tk.Pair = pair

	return tk, responseBody, err
}

func (okx *OKxV5) GetKline(pair CurrencyPair, period KlinePeriod, opt ...OptionParameter) ([]Kline, []byte, error) {
	reqUrl := fmt.Sprintf("%s%s", okx.UriOpts.Endpoint, okx.UriOpts.KlineUri)
	param := url.Values{}
	param.Set("instId", pair.Symbol)
	param.Set("bar", AdaptKlinePeriodToSymbol(period))
	param.Set("limit", "100")
	MergeOptionParams(&param, opt...)

	data, responseBody, err := okx.DoNoAuthRequest(http.MethodGet, reqUrl, &param)
	if err != nil {
		return nil, nil, err
	}
	klines, err := okx.UnmarshalOpts.KlineUnmarshaler(data)
	return klines, responseBody, err
}

func (okx *OKxV5) GetHistoryKline(pair CurrencyPair, period KlinePeriod, opt ...OptionParameter) ([]Kline, []byte, error) {
	reqUrl := fmt.Sprintf("%s%s", okx.UriOpts.Endpoint, okx.UriOpts.HistoryKlineUri)
	param := url.Values{}
	param.Set("instId", pair.Symbol)
	param.Set("bar", AdaptKlinePeriodToSymbol(period))
	param.Set("limit", "100")
	MergeOptionParams(&param, opt...)

	data, responseBody, err := okx.DoNoAuthRequest(http.MethodGet, reqUrl, &param)
	if err != nil {
		return nil, nil, err
	}
	klines, err := okx.UnmarshalOpts.KlineUnmarshaler(data)
	return klines, responseBody, err
}

func (okx *OKxV5) GetExchangeInfo(instType string, opt ...OptionParameter) (map[string]CurrencyPair, []byte, error) {
	reqUrl := fmt.Sprintf("%s%s", okx.UriOpts.Endpoint, okx.UriOpts.GetExchangeInfoUri)
	param := url.Values{}
	param.Set("instType", instType)
	MergeOptionParams(&param, opt...)

	data, responseBody, err := okx.DoNoAuthRequest(http.MethodGet, reqUrl, &param)
	if err != nil {
		return nil, responseBody, err
	}

	currencyPairMap, err := okx.UnmarshalOpts.GetExchangeInfoResponseUnmarshaler(data)

	return currencyPairMap, responseBody, err
}

func (okx *OKxV5) DoNoAuthRequest(httpMethod, reqUrl string, params *url.Values) ([]byte, []byte, error) {
	reqBody := ""
	if http.MethodGet == httpMethod {
		reqUrl += "?" + params.Encode()
	}

	responseBody, err := Cli.DoRequest(httpMethod, reqUrl, reqBody, nil)
	if err != nil {
		return nil, responseBody, err
	}

	var baseResp BaseResp
	err = okx.UnmarshalOpts.ResponseUnmarshaler(responseBody, &baseResp)
	if err != nil {
		return responseBody, responseBody, err
	}

	if baseResp.Code == 0 {
		logger.Debugf("[DoNoAuthRequest] response=%s", string(responseBody))
		return baseResp.Data, responseBody, nil
	}

	logger.Debugf("[DoNoAuthRequest] error=%s", baseResp.Msg)
	return nil, responseBody, err
}
