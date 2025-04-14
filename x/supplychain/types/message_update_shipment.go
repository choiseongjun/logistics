package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateShipment{}

func NewMsgUpdateShipment(creator string, shipmentId string, newStatus string, newLocation string) *MsgUpdateShipment {
	return &MsgUpdateShipment{
		Creator:     creator,
		ShipmentId:  shipmentId,
		NewStatus:   newStatus,
		NewLocation: newLocation,
	}
}

func (msg *MsgUpdateShipment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
