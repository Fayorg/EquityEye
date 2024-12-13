package provider

import (
	"EquityEye/internal/logs"
	"EquityEye/types"
	"errors"
)

type Provider interface {
	GetProviderName() string
	GetProviderConfigName() string

	GetAvailableTicker() []types.Ticker
	GetMarketDataForTicker(ticker types.Ticker) (float64, error)
}

var providers = make(map[string]Provider)
var availableTickers = make(map[types.Ticker][]Provider)

func InitializeProviders(provider []types.ProviderConfiguration) {
	for _, p := range provider {
		switch p.ProviderName {
		case "BINANCE":
			binanceProvider := NewBinanceProvider(p)
			providers[p.ProviderName] = binanceProvider
			for _, ticker := range binanceProvider.GetAvailableTicker() {
				availableTickers[ticker] = append(availableTickers[ticker], binanceProvider)
			}
			continue
		default:
			logs.Info("Provider %s not supported", p.ProviderName)
			continue
		}
	}
}

func GetGloballyAvailableTickers() []types.Ticker {
	var tickers []types.Ticker
	for ticker := range availableTickers {
		tickers = append(tickers, ticker)
	}
	return tickers
}

func GetProviderForTicker(ticker types.Ticker) ([]Provider, error) {
	providers := availableTickers[ticker]
	if len(providers) == 0 {
		return nil, errors.New("no provider could be found for this ticker")
	}
	return providers, nil
}
