package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/vincadian/arkh-blockchain/testutil/keeper"
	"github.com/vincadian/arkh-blockchain/x/tool/keeper"
	"github.com/vincadian/arkh-blockchain/x/tool/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ToolKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
