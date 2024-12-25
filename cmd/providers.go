package cmd

import (
	"EquityEye/internal/cache"
	"EquityEye/internal/config"
	"EquityEye/internal/logs"
	"EquityEye/internal/provider"
	"github.com/spf13/cobra"
)

var providersCmd = &cobra.Command{
	Use:   "providers",
	Short: "Initialize providers and make operations on them",
	RunE:  runProviders,
}

var tickersCmd = &cobra.Command{
	Use:   "tickers",
	Short: "Get tickers from providers",
	RunE:  runTickers,
}

func init() {
	// Config Subcommand
	providersCmd.AddCommand(tickersCmd)
}

func runProviders(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	logs.Info("Configuration is valid")

	cache := cache.NewRedisCache(cfg.Cache.Url)
	provider.InitializeProviders(cfg.Providers, cache)

	logs.Info("Providers initialized")
	return nil
}

func runTickers(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	logs.Info("Configuration is valid")

	cache := cache.NewRedisCache(cfg.Cache.Url)
	provider.InitializeProviders(cfg.Providers, cache)
	logs.Info("Providers initialized")

	tickers := provider.GetAllTickers()

	logs.Info("List of available tickers from configured providers")
	for _, ticker := range tickers {
		logs.Info("%v", ticker)
	}

	logs.Info("Total number of tickers available: %v", len(tickers))
	return nil
}
