package provider

import (
	"EquityEye/internal/cache"
	"EquityEye/internal/logs"
	"EquityEye/types"
	"errors"
)

type Provider interface {
	GetProviderName() string
	GetProviderConfiguration() types.ProviderConfiguration

	InitializeProvider() error

	GetAvailableTickers() []types.Ticker
	GetMarketDataForTicker(ticker types.Ticker) (float64, error)
}

// TODO: find a more elegant/memory efficient way to cache this
var providers = make(map[string]Provider)
var availableTickers = make(map[types.Ticker][]Provider)
var tickersCache = make(map[string]types.Ticker)

func InitializeProviders(provider []types.ProviderConfiguration, cache cache.Cache) {
	for _, p := range provider {
		err := cache.RegisterProvider(p)
		if err != nil {
			logs.Error("Could not register provider %s", p.ProviderName)
			continue
		}

		var provider Provider
		switch p.ProviderName {
		case "BINANCE":
			provider = NewBinanceProvider(cache, p)
		default:
			logs.Info("Provider %s not supported", p.ProviderName)
			continue
		}
		err = provider.InitializeProvider()
		if err != nil {
			logs.Error("Could not initialize provider %s", p.ProviderName)
			continue
		}
		providers[p.ProviderName] = provider
		for _, ticker := range provider.GetAvailableTickers() {
			availableTickers[ticker] = append(availableTickers[ticker], provider)
		}
	}
}

func GetAllTickers() []types.Ticker {
	if len(tickersCache) > 0 {
		var tickers []types.Ticker
		for _, ticker := range tickersCache {
			tickers = append(tickers, ticker)
		}
		return tickers
	}
	var tickers []types.Ticker
	for ticker := range availableTickers {
		tickers = append(tickers, ticker)
	}
	return tickers
}

func GetTicker(ticker string) (types.Ticker, error) {
	if ticker, ok := tickersCache[ticker]; ok {
		return ticker, nil
	}
	tickers := GetAllTickers()
	for _, t := range tickers {
		if t.Tick == ticker {
			return t, nil
		}
	}
	return types.Ticker{}, errors.New("ticker not found")
}

func GetProvidersForTicker(ticker types.Ticker) ([]Provider, error) {
	providers := availableTickers[ticker]
	if len(providers) == 0 {
		return nil, errors.New("no provider could be found for this ticker")
	}
	return providers, nil
}
