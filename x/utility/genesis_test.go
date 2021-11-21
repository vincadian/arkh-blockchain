package utility_test

import (
	"testing"

	keepertest "github.com/arkhadian/arkh/testutil/keeper"
	"github.com/arkhadian/arkh/x/utility"
	"github.com/arkhadian/arkh/x/utility/types"
	"github.com/stretchr/testify/require"
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
