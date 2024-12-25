package main

import (
	"EquityEye/cmd"
	"EquityEye/internal/logs"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logs.Error(err.Error())
		os.Exit(1)
	}
}
