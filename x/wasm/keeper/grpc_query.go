package keeper

import (
	"github.com/arkhadian/arkh/x/wasm/types"
)

var _ types.QueryServer = Keeper{}
