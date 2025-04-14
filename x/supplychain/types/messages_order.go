package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateOrder{}

func NewMsgCreateOrder(creator string, orderId string, purchaseId string, seller string, buyer string, status string, items string, price string, timestamp string) *MsgCreateOrder {
	return &MsgCreateOrder{
		Creator:    creator,
		OrderId:    orderId,
		PurchaseId: purchaseId,
		Seller:     seller,
		Buyer:      buyer,
		Status:     status,
		Items:      items,
		Price:      price,
		Timestamp:  timestamp,
	}
}

func (msg *MsgCreateOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateOrder{}

func NewMsgUpdateOrder(creator string, id uint64, orderId string, purchaseId string, seller string, buyer string, status string, items string, price string, timestamp string) *MsgUpdateOrder {
	return &MsgUpdateOrder{
		Id:         id,
		Creator:    creator,
		OrderId:    orderId,
		PurchaseId: purchaseId,
		Seller:     seller,
		Buyer:      buyer,
		Status:     status,
		Items:      items,
		Price:      price,
		Timestamp:  timestamp,
	}
}

func (msg *MsgUpdateOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteOrder{}

func NewMsgDeleteOrder(creator string, id uint64) *MsgDeleteOrder {
	return &MsgDeleteOrder{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
