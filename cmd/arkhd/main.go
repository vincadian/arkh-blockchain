package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/vincadian/arkh-blockchain/app"
)

func main() {
	rootCmd, _ := app.NewRootCmd()
	if err := server.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
