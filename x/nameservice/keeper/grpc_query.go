package keeper

import (
	"github.com/vincadian/arkh-blockchain/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
