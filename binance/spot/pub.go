package spot

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	. "github.com/wfunc/goex/httpcli"
	"github.com/wfunc/goex/logger"
	. "github.com/wfunc/goex/model"
	. "github.com/wfunc/goex/util"
)

func (s *Spot) GetName() string {
	return "binance.com"
}

func (s *Spot) GetDepth(pair CurrencyPair, size int, opts ...OptionParameter) (*Depth, []byte, error) {
	params := url.Values{}
	params.Set("symbol", pair.Symbol)
	params.Set("limit", fmt.Sprint(size))
	MergeOptionParams(&params, opts...)

	reqUrl := fmt.Sprintf("%s%s", s.UriOpts.Endpoint, s.UriOpts.DepthUri)
	data, err := s.DoNoAuthRequest(http.MethodGet, reqUrl, &params, nil)
	if err != nil {
		return nil, data, err
	}
	logger.Debugf("[GetDepth] %s", string(data))
	dep, err := s.UnmarshalerOpts.DepthUnmarshaler(data)
	return dep, data, err
}

func (s *Spot) GetTicker(pair CurrencyPair, opt ...OptionParameter) (*Ticker, []byte, error) {
	params := url.Values{}
	params.Set("symbol", pair.Symbol)

	if len(opt) > 0 {
		for _, p := range opt {
			if p.Key == "symbols" {
				params.Del("symbol") //only symbol or symbols
			}
			params.Add(p.Key, p.Value)
		}
	}

	data, err := s.DoNoAuthRequest(http.MethodGet,
		fmt.Sprintf("%s%s", s.UriOpts.Endpoint, s.UriOpts.TickerUri), &params, nil)
	if err != nil {
		return nil, data, fmt.Errorf("%w%s", err, errors.New(string(data)))
	}

	tk, err := s.UnmarshalerOpts.TickerUnmarshaler(data)
	if err != nil {
		return nil, data, err
	}

	tk.Pair = pair

	return tk, data, err
}

func (s *Spot) GetKline(pair CurrencyPair, period KlinePeriod, opts ...OptionParameter) ([]Kline, []byte, error) {
	params := url.Values{}
	params.Set("limit", "1000")
	params.Set("symbol", pair.Symbol)
	params.Set("interval", adaptKlinePeriod(period))

	MergeOptionParams(&params, opts...)

	reqUrl := fmt.Sprintf("%s%s", s.UriOpts.Endpoint, s.UriOpts.KlineUri)
	respBody, err := s.DoNoAuthRequest(http.MethodGet, reqUrl, &params, nil)
	if err != nil {
		return nil, respBody, err
	}

	klines, err := s.UnmarshalerOpts.KlineUnmarshaler(respBody)
	return klines, respBody, err
}

func (s *Spot) GetExchangeInfo() (map[string]CurrencyPair, []byte, error) {
	panic("not implement")
}

func (s *Spot) DoNoAuthRequest(method, reqUrl string, params *url.Values, headers map[string]string) ([]byte, error) {
	var reqBody string

	if method == http.MethodGet {
		reqUrl += "?" + params.Encode()
	} else {
		reqBody = params.Encode()
	}

	responseData, err := Cli.DoRequest(method, reqUrl, reqBody, headers)
	if err != nil {
		return responseData, err
	}

	return responseData, err
}
