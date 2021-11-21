package nameservice_test

import (
	"testing"

	keepertest "github.com/arkhadian/arkh/testutil/keeper"
	"github.com/arkhadian/arkh/x/nameservice"
	"github.com/arkhadian/arkh/x/nameservice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NameserviceKeeper(t)
	nameservice.InitGenesis(ctx, *k, genesisState)
	got := nameservice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
