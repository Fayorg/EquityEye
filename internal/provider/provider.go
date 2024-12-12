package provider

import "EquityEye/types"

type Provider interface {
	GetAvailableTicker() []types.Ticker
	GetMarketDataForTicker(ticker types.Ticker) (float64, error)
}

var providersHashMap = make(map[string]Provider)

func InitializeProviders(provider []types.Provider) {

}
