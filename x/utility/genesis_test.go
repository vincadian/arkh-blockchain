package utility_test

import (
	"testing"

	keepertest "github.com/vincadian/arkh-blockchain/testutil/keeper"
	"github.com/stretchr/testify/require"
	"github.com/vincadian/arkh-blockchain/x/utility"
	"github.com/vincadian/arkh-blockchain/x/utility/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UtilityKeeper(t)
	utility.InitGenesis(ctx, *k, genesisState)
	got := utility.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
