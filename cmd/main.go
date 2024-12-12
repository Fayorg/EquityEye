package main

import (
	cache "EquityEye/internal/cache"
	config "EquityEye/internal/config"
	"EquityEye/internal/logs"
	"fmt"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		// logs.Error("Could not load config")
		logs.Error(err.Error())
		os.Exit(0)
	}
	cache := cache.NewRedisCache(cfg.Cache.Url)

	fmt.Println("Using config : ", cfg.Environment)

	fmt.Println("Hello World")
}
