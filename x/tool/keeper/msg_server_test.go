package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arkhadian/arkh/testutil/keeper"
	"github.com/arkhadian/arkh/x/tool/keeper"
	"github.com/arkhadian/arkh/x/tool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ToolKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
