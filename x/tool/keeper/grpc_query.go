package keeper

import (
	"github.com/vincadian/arkh-blockchain/x/tool/types"
)

var _ types.QueryServer = Keeper{}
