package main

import (
	cache "EquityEye/internal/cache"
	config "EquityEye/internal/config"
	"EquityEye/internal/logs"
	"EquityEye/internal/provider"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		// logs.Error("Could not load config")
		logs.Error(err.Error())
		os.Exit(0)
	}
	_ = cache.NewRedisCache(cfg.Cache.Url)

	provider.InitializeProviders(cfg.Providers)

	tickers := provider.GetGloballyAvailableTickers()
	logs.Info("Available tickers: %v", tickers)

	providers, err := provider.GetProviderForTicker(tickers[0])
	logs.Info("Available provider for ticker %v: ", tickers[0])
	for _, p := range providers {
		logs.Info("%v", p.GetProviderConfigName())
	}
}
