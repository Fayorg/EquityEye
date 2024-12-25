package cmd

import (
	"EquityEye/cmd/api"
	"EquityEye/internal/logs"
	"context"
	"github.com/spf13/cobra"
	"os"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the API server",
	RunE:  runApi,
}

func init() {
	// Config Subcommand
}

func runApi(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	if err := api.Run(ctx, os.Stdout, os.Args); err != nil {
		logs.Error("%s\n", err)
		return nil
	}
	return nil
}
