package incentive

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lcnem/eurx/x/incentive/keeper"
	"github.com/lcnem/eurx/x/incentive/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgClaimEurxMintingReward:
			return handleMsgClaimEurxMintingReward(ctx, k, msg)
		// this line is used by starport scaffolding # 1
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgClaimEurxMintingReward(ctx sdk.Context, k keeper.Keeper, msg *types.MsgClaimEurxMintingReward) (*sdk.Result, error) {

	err := k.ClaimEurxMintingReward(ctx, msg.Sender.AccAddress(), msg.MultiplierName)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}
