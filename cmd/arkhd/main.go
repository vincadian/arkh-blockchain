package main

import (
	"os"

	"github.com/vincadian/arkh-blockchain/app"
)

func main() {
	rootCmd, err := app.NewRootCmd()
	if err != nil {
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
