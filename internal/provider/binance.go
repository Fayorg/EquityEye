package provider

import (
	"EquityEye/internal/cache"
	"EquityEye/internal/logs"
	"EquityEye/types"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type BinanceProvider struct {
	Cache         cache.Cache
	Configuration types.ProviderConfiguration
}

const (
	BinanceProviderName = "BINANCE"
	BinanceAPIURL       = "https://api.binance.com"
)

var BinanceTickers = map[types.Ticker]string{
	{Name: "BTC", Tick: "BTC", Market: "USD"}: "BTCUSDT",
}

func NewBinanceProvider(cache cache.Cache, configuration types.ProviderConfiguration) Provider {
	return &BinanceProvider{
		Cache:         cache,
		Configuration: configuration,
	}
}

func (b *BinanceProvider) GetProviderName() string {
	return BinanceProviderName
}

func (b *BinanceProvider) GetProviderConfigName() string {
	return b.Configuration.Name
}

func (b *BinanceProvider) GetProviderConfiguration() types.ProviderConfiguration {
	return b.Configuration
}

func (b *BinanceProvider) GetAvailableTicker() []types.Ticker {
	var tickers []types.Ticker
	for ticker := range BinanceTickers {
		tickers = append(tickers, ticker)
	}
	return tickers
}

func (b *BinanceProvider) GetMarketDataForTicker(ticker types.Ticker) (float64, error) {
	type ResponseData struct {
		Price  float64 `json:"price,string"`
		Symbol string  `json:"symbol"`
	}
	dis, err := b.Cache.IsProviderTemporarilyDisabled(b.Configuration)
	if err != nil {
		return -1, err
	}
	if dis {
		return -1, errors.New("provider is temporarily disabled")
	}

	resp, err := http.Get(fmt.Sprintf("%s/api/v3/ticker/price?symbol=%s", BinanceAPIURL, BinanceTickers[ticker]))
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	if err := b.checkForRateLimit(resp); err != nil {
		return -1, err
	}
	/*if err := b.updateUsage(resp); err != nil {
		return -1, err
	}*/

	var data ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		logs.Error("Failed to parse JSON: %v", err)
	}

	return data.Price, nil
}

func (b *BinanceProvider) checkForRateLimit(res *http.Response) error {
	if res.StatusCode != 418 && res.StatusCode != 429 {
		return nil
	}

	wait, err := strconv.Atoi(res.Header.Get("Retry-After"))
	if err != nil {
		logs.Warn("Couldn't parse Retry-After header for provider %s. This could result in a IP-ban", b.Configuration.Name)
		return errors.New("couldn't parse Retry-After header")
	}

	err = b.Cache.TemporaryDisableProvider(b.Configuration, time.Duration(wait)*time.Second)
	if err != nil {
		logs.Warn("Couldn't disable provider %s. This could result in a IP-ban", b.Configuration.Name)
	}

	if res.StatusCode == 429 {
		return errors.New("rate limit exceeded")
	}
	if res.StatusCode == 418 {
		logs.Warn("IP has been auto-banned due to too many requests")

		return errors.New("ip has been auto-banned due to too many requests")
	}

	return errors.New("unknown error")
}

func (b *BinanceProvider) updateUsage(res *http.Response) error {
	weight, err := strconv.Atoi(res.Header.Get("X-MBX-USED-WEIGHT"))
	if err != nil {
		logs.Warn("Couldn't parse X-MBX-USED-WEIGHT header for provider %s", b.Configuration.Name)
		return errors.New("couldn't parse X-MBX-USED-WEIGHT header")
	}

	return b.Cache.IncreaseUsageBy(b.Configuration, weight)
}
