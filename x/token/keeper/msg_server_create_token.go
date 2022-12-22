package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/platine-network/platine/x/token/types"
)

func (k msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenID := "u" + strings.ToLower(msg.Symbol)
	err := k.CheckCommonError(tokenID, msg.Symbol, msg.Name, msg.Supply)
	if err != nil {
		return nil, err
	}

	_, found := k.GetToken(ctx, tokenID)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token already exists")
	}

	ownerAddress, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	const mintSequence = 1
	const burnSequence = 0

	var token = types.Token{
		Id:           tokenID,
		Owner:        msg.Owner,
		Name:         msg.Name,
		Symbol:       msg.Symbol,
		Supply:       msg.Supply,
		Decimal:      msg.Decimal,
		Mintable:     msg.Mintable,
		Burnable:     msg.Burnable,
		MintSequence: mintSequence,
		BurnSequence: burnSequence,
	}

	k.SetToken(ctx, token)

	newCoin := sdk.NewInt64Coin(tokenID, int64(msg.Supply))

	err = k.bankKeeper.MintCoins(
		ctx,
		types.ModuleName,
		sdk.NewCoins(newCoin),
	)

	if err != nil {
		return nil, err
	}

	// transfer to account
	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		ownerAddress,
		sdk.NewCoins(newCoin),
	)

	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute("Owner", msg.Owner),
		sdk.NewAttribute("TokenID", tokenID),
		sdk.NewAttribute("TokenName", msg.Name),
		sdk.NewAttribute("TokenSymbol", msg.Symbol),
	))

	return &types.MsgCreateTokenResponse{}, nil
}
