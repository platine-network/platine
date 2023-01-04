package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/platine-network/platine/x/token/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token, found := k.GetToken(ctx, msg.TokenID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "token not found")
	}

	if !token.Mintable {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "mint is not allowed")
	}

	err := k.CheckCommonError(token.Id, token.Symbol, token.Name, token.Supply)
	if err != nil {
		return nil, err
	}

	if msg.Owner != token.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only owner can process the burn")
	}

	newSupply := msg.Amount + token.Supply
	if newSupply > maxTokenSupply {
		errMsg := fmt.Sprintf("token total supply can not exceed %d", maxTokenSupply)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, errMsg)
	}

	newCoin := sdk.NewInt64Coin(msg.TokenID, int64(msg.Amount))

	err = k.bankKeeper.MintCoins(
		ctx,
		types.ModuleName,
		sdk.NewCoins(newCoin),
	)

	if err != nil {
		return nil, err
	}

	receiverAddress, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	// transfer to account
	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		receiverAddress,
		sdk.NewCoins(newCoin),
	)

	if err != nil {
		return nil, err
	}

	token.MintSequence++
	token.Supply += msg.Amount

	k.SetToken(ctx, token)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenID", msg.TokenID),
		sdk.NewAttribute("MintAmount", strconv.FormatUint(msg.Amount, 10)),
		sdk.NewAttribute("Supply", strconv.FormatUint(token.Supply, 10)),
	))

	return &types.MsgMintResponse{}, nil
}
