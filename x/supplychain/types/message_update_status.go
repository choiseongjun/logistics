package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateStatus{}

func NewMsgUpdateStatus(creator string, productId string, newStatus string) *MsgUpdateStatus {
	return &MsgUpdateStatus{
		Creator:   creator,
		ProductId: productId,
		NewStatus: newStatus,
	}
}

func (msg *MsgUpdateStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
