package keeper

import (
	"github.com/vincadian/arkh-blockchain/x/utility/types"
)

var _ types.QueryServer = Keeper{}
