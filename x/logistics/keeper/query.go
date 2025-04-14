package keeper

import (
	"logistics/x/logistics/types"
)

var _ types.QueryServer = Keeper{}
