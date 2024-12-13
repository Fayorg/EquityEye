package provider

import "EquityEye/types"

type BinanceProvider struct {
	ConfigName string
	Weight     int
	Key        string
}

const BinanceProviderName = "BINANCE"

var BinanceTickers = map[types.Ticker]string{
	{Name: "BTC", Tick: "BTC", Market: "USD"}: "BTCUSD",
}

func NewBinanceProvider(configuration types.ProviderConfiguration) Provider {
	return &BinanceProvider{
		ConfigName: configuration.Name,
		Weight:     configuration.Weight,
		Key:        configuration.Key,
	}
}

func (b *BinanceProvider) GetProviderName() string {
	return BinanceProviderName
}

func (b *BinanceProvider) GetProviderConfigName() string {
	return b.ConfigName
}

func (b *BinanceProvider) GetAvailableTicker() []types.Ticker {
	var tickers []types.Ticker
	for ticker := range BinanceTickers {
		tickers = append(tickers, ticker)
	}
	return tickers
}

func (b *BinanceProvider) GetMarketDataForTicker(ticker types.Ticker) (float64, error) {
	return 0, nil
}
