package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MintHooks interface {
	AfterDistributeMintedCoin(ctx sdk.Context, mintedCoin sdk.Coin)
}
type MultiMintHooks []MintHooks

var _ MintHooks = MultiMintHooks{}

func NewMultiMintHooks(hooks ...MintHooks) MultiMintHooks {
	return hooks
}

func (h MultiMintHooks) AfterDistributeMintedCoin(ctx sdk.Context, mintedCoin sdk.Coin) {
	for i := range h {
		h[i].AfterDistributeMintedCoin(ctx, mintedCoin)
	}
}

