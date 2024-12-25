package cmd

import (
	"EquityEye/internal/config"
	"EquityEye/internal/logs"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Check the configuration of the application and detect any issues",
	RunE:  runConfig,
}

func init() {
	// Config Subcommand
}

func runConfig(cmd *cobra.Command, args []string) error {
	_, err := config.LoadConfig()
	if err != nil {
		logs.Error(err.Error())
		return nil
	}

	logs.Info("Configuration is valid")
	return nil
}
