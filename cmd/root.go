package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "equityeye",
	Short: "Market price monitoring tool for cryptocurrencies, stocks and much more",
	Long:  `EquityEye is a market price monitoring tool for cryptocurrencies, stocks and much more. It provides a simple interface to monitor market prices from multiple providers and store it in a database for future use..`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add sub-commands

	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(providersCmd)
	// rootCmd.AddCommand(produceCmd)
}
