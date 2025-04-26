package keeper

import (
	"github.com/vincadian/arkh-blockchain/x/wasm/types"
)

var _ types.QueryServer = Keeper{}
