package keeper

import (
	"github.com/lcnem/eurx/x/pricefeed/types"
)

var _ types.QueryServer = Keeper{}
