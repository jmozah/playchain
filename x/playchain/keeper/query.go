package keeper

import (
	"playchain/x/playchain/types"
)

var _ types.QueryServer = Keeper{}
