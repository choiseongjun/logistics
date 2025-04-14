package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferOwnership{}

func NewMsgTransferOwnership(creator string, productId string, newOwner string, location string) *MsgTransferOwnership {
	return &MsgTransferOwnership{
		Creator:   creator,
		ProductId: productId,
		NewOwner:  newOwner,
		Location:  location,
	}
}

func (msg *MsgTransferOwnership) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
