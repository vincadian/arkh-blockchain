package simapp

import (
	// "cosmossdk.io/simapp" // Removed to avoid dependency conflicts
	// Note: tendermint/spm/cosmoscmd deprecated in Cosmos SDK v0.53, now using CometBFT
	cosmoslog "cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	tmdb "github.com/cosmos/cosmos-db"

	"github.com/vincadian/arkh-blockchain/app"
)

// EmptyAppOptions is a stub implementing servertypes.AppOptions with empty values.
type EmptyAppOptions struct{}

// Get implements servertypes.AppOptions
func (ao EmptyAppOptions) Get(string) interface{} {
	return nil
}

// New creates application instance with in-memory database and disabled logging.
func New(dir string) *app.App {
	db := tmdb.NewMemDB()
	logger := cosmoslog.NewNopLogger()

	encoding := app.MakeEncodingConfig()

	a := app.New(logger, db, nil, true, map[int64]bool{}, dir, 0, encoding,
		EmptyAppOptions{})
	// InitChain updates deliverState which is required when app.NewContext is called
	a.InitChain(&abci.RequestInitChain{
		AppStateBytes: []byte("{}"),
	})
	return a
}
