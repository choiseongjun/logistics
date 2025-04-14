package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateProduct{}

func NewMsgCreateProduct(creator string, productId string, name string, origin string, currentOwner string, status string) *MsgCreateProduct {
	return &MsgCreateProduct{
		Creator:      creator,
		ProductId:    productId,
		Name:         name,
		Origin:       origin,
		CurrentOwner: currentOwner,
		Status:       status,
	}
}

func (msg *MsgCreateProduct) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
