package keeper

import (
	"github.com/arkhadian/arkh/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
