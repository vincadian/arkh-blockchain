package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/vincadian/arkh-blockchain/app"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "panic: %v\n\n", r)
			debug.PrintStack()
			os.Exit(1)
		}
	}()

	rootCmd, err := app.NewRootCmd()
	if err != nil {
		os.Exit(1)
	}

	// Set server and client context in command context so start command and gRPC/API (e.g. SignModeHandler) work.
	ctx := context.WithValue(context.Background(), server.ServerContextKey, &server.Context{})
	ctx = context.WithValue(ctx, client.ClientContextKey, &client.Context{})
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
