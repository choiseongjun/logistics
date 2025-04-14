package keeper

import (
	"logistics/x/supplychain/types"
)

var _ types.QueryServer = Keeper{}
