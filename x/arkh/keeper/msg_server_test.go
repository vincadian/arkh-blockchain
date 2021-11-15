package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arkhadian/arkh/testutil/keeper"
	"github.com/arkhadian/arkh/x/arkh/keeper"
	"github.com/arkhadian/arkh/x/arkh/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ArkhKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
