package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arkhadian/arkh/testutil/keeper"
	"github.com/arkhadian/arkh/x/utility/keeper"
	"github.com/arkhadian/arkh/x/utility/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.UtilityKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
