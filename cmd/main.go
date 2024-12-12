package main

import (
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

	fmt.Println("Using config : ", cfg.Environment)

	fmt.Println("Hello World")
}
