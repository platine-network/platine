package keeper

import (
	"context"
	"strconv"

	"github.com/platine-network/platine/x/token/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token, found := k.GetToken(ctx, msg.TokenID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "token not found")
	}

	if !token.Burnable {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "burn is not allowed")
	}

	err := k.CheckCommonError(token.Id, token.Symbol, token.Name, token.Supply)
	if err != nil {
		return nil, err
	}

	if msg.Owner != token.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only owner can process the burn")
	}

	ownerAddress, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	balance := k.bankKeeper.GetBalance(ctx, ownerAddress, msg.TokenID).Amount
	if balance.Uint64() < msg.Amount {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "account balance is not enough to burn")
	}

	newCoin := sdk.NewInt64Coin(msg.TokenID, int64(msg.Amount))

	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		ownerAddress,
		types.ModuleName,
		sdk.NewCoins(newCoin),
	)

	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(
		ctx,
		types.ModuleName,
		sdk.NewCoins(newCoin),
	)

	if err != nil {
		return nil, err
	}

	token.BurnSequence++
	token.Supply -= msg.Amount

	k.SetToken(ctx, token)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenID", msg.TokenID),
		sdk.NewAttribute("BurnAmount", strconv.FormatUint(msg.Amount, 10)),
		sdk.NewAttribute("Supply", strconv.FormatUint(token.Supply, 10)),
	))

	return &types.MsgBurnResponse{}, nil
}
