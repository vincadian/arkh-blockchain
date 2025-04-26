package keeper

import (
	"github.com/vincadian/arkh-blockchain/x/arkh/types"
)

var _ types.QueryServer = Keeper{}
