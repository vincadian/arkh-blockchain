package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/arkh module sentinel errors
var (
	ErrSample = errors.Register(ModuleName, 1100, "sample error")
)
