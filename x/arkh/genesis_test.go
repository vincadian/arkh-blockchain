package arkh_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vincadian/arkh-blockchain/testutil/keeper"
	"github.com/vincadian/arkh-blockchain/x/arkh"
	"github.com/vincadian/arkh-blockchain/x/arkh/types"
)

func TestGenesis(t *testing.T) {
	// Initialize a meaningful genesis state
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		// Add other fields as needed
	}

	// Initialize the keeper and context
	k, ctx := keepertest.ArkhKeeper(t)

	// Initialize the genesis state
	arkh.InitGenesis(ctx, *k, genesisState)

	// Export the genesis state
	got := arkh.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// Validate the exported genesis state
	require.Equal(t, genesisState.Params, got.Params)
	// Add more assertions as needed
}
