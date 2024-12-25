package main

import (
	cache "EquityEye/internal/cache"
	config "EquityEye/internal/config"
	"EquityEye/internal/logs"
	"EquityEye/internal/provider"
	"os"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		// logs.Error("Could not load config")
		logs.Error(err.Error())
		os.Exit(0)
	}
	cache := cache.NewRedisCache(cfg.Cache.Url)

	provider.InitializeProviders(cfg.Providers, cache)

	ticker, err := provider.GetTicker("ETH")
	if err != nil {
		logs.Error("Could not get ticker %v", "ETH")
		logs.Error(err.Error())
		os.Exit(0)
	}

	providers, err := provider.GetProvidersForTicker(ticker)
	logs.Info("Available provider for ticker %v: ", ticker)
	for _, p := range providers {
		logs.Info("%v", p.GetProviderConfiguration().Name)
	}

	// for loop that runs 10 times
	for i := 0; i < 2; i++ {
		btc, err := providers[0].GetMarketDataForTicker(ticker)
		if err != nil {
			logs.Error("Could not get market data for ticker %v", ticker)
			logs.Error(err.Error())
		} else {
			logs.Info("Market data for ticker %v: %v", ticker, btc)
		}
		time.Sleep(1 * time.Second)
	}

	val, err := cache.GetUsage(providers[0].GetProviderConfiguration())
	if err != nil {
		logs.Warn("Could not get usage for provider %v", providers[0].GetProviderName())
	}
	logs.Info("Usage for provider %v: %d", providers[0].GetProviderName(), val)

	/*time.Sleep(5 * time.Second)

	err = cache.TemporaryDisableProvider(providers[0].GetProviderConfiguration(), 10*time.Second)
	if err != nil {
		logs.Warn("Could not disable provider %v", providers[0].GetProviderName())
	} else {
		logs.Info("Provider %v disabled for 10 seconds", providers[0].GetProviderName())
	}

	btc, err = providers[0].GetMarketDataForTicker(tickers[0])
	if err != nil {
		logs.Error("Could not get market data for ticker %v", tickers[0])
		logs.Error(err.Error())
	} else {
		logs.Info("Market data for ticker %v: %v", tickers[0], btc)
	}*/
}
