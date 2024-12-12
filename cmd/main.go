package main

import (
	cache "EquityEye/internal/cache"
	config "EquityEye/internal/config"
	"EquityEye/internal/logs"
	"fmt"
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

	for _, provider := range cfg.Providers {
		err := cache.RegisterProvider(provider)
		if err != nil {
			logs.Error("Could not register provider %s", provider.Name)
			os.Exit(0)
		}
	}

	for true {
		err := cache.RegisterProvider(cfg.Providers[0])
		if err != nil {
			logs.Error("Could not register provider %s", cfg.Providers[0].Name)
			os.Exit(0)
		}
		val, err := cache.GetUsage(cfg.Providers[0])
		if err != nil {
			logs.Error("Could not get usage for provider %s", cfg.Providers[0].Name)
			os.Exit(0)
		}
		logs.Info("Usage for provider %s is %d", cfg.Providers[0].Name, val)

		// Sleep for 1 second
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Using config : ", cfg.Environment)

	fmt.Println("Hello World")
}
