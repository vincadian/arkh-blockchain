package wasm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vincadian/arkh-blockchain/testutil/keeper"
	"github.com/vincadian/arkh-blockchain/x/wasm"
	"github.com/vincadian/arkh-blockchain/x/wasm/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WasmKeeper(t)
	wasm.InitGenesis(ctx, *k, genesisState)
	got := wasm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
