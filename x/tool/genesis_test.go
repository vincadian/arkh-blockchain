package tool_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vincadian/arkh-blockchain/testutil/keeper"
	"github.com/vincadian/arkh-blockchain/x/tool"
	"github.com/vincadian/arkh-blockchain/x/tool/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ToolKeeper(t)
	tool.InitGenesis(ctx, *k, genesisState)
	got := tool.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
