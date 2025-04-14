package keeper

import (
	"context"
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"logistics/x/supplychain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateShipment(goCtx context.Context, msg *types.MsgUpdateShipment) (*types.MsgUpdateShipmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	shipmentId, err := strconv.ParseUint(msg.ShipmentId, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid product ID format")
	}
	// 배송 정보 존재 여부 확인
	shipment, found := k.GetShipment(ctx, shipmentId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "shipment not found")
	}

	// 권한 확인 (sender 또는 receiver만 업데이트 가능)
	if msg.Creator != shipment.Sender && msg.Creator != shipment.Receiver {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "unauthorized: only sender or receiver can update")
	}

	// 타임스탬프 생성
	timestamp := time.Now().Unix()

	// 배송 정보 업데이트
	shipment.Status = msg.NewStatus
	shipment.Location = msg.NewLocation
	shipment.Timestamp = strconv.FormatInt(timestamp, 10)

	productId, err := strconv.ParseUint(shipment.ProductId, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid product ID format")
	}

	// 배송 정보 저장
	k.SetShipment(ctx, shipment)

	// 배송이 완료되면 제품 상태도 자동 업데이트 (스마트 계약 기능)
	if msg.NewStatus == "delivered" {
		product, found := k.GetProduct(ctx, productId)
		if found {
			product.Status = "delivered"
			// 소유권도 receiver로 변경
			product.CurrentOwner = shipment.Receiver
			k.SetProduct(ctx, product)

			// 소유권 이전 이벤트 발생
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeOwnershipTransferred,
					sdk.NewAttribute(types.AttributeKeyProductId, shipment.ProductId),
					sdk.NewAttribute(types.AttributeKeyPreviousOwner, shipment.Sender),
					sdk.NewAttribute(types.AttributeKeyNewOwner, shipment.Receiver),
					sdk.NewAttribute(types.AttributeKeyTimestamp, strconv.FormatInt(timestamp, 10)),
					sdk.NewAttribute(types.AttributeKeyLocation, msg.NewLocation),
				),
			)
		}
	}

	// 배송 업데이트 이벤트 발생
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeShipmentUpdated,
			sdk.NewAttribute(types.AttributeKeyShipmentId, msg.ShipmentId),
			sdk.NewAttribute(types.AttributeKeyProductId, shipment.ProductId),
			sdk.NewAttribute(types.AttributeKeyOldStatus, shipment.Status),
			sdk.NewAttribute(types.AttributeKeyNewStatus, msg.NewStatus),
			sdk.NewAttribute(types.AttributeKeyLocation, msg.NewLocation),
			sdk.NewAttribute(types.AttributeKeyTimestamp, strconv.FormatInt(timestamp, 10)),
		),
	)

	return &types.MsgUpdateShipmentResponse{
		ShipmentId: msg.ShipmentId,
		Timestamp:  strconv.FormatInt(timestamp, 10),
	}, nil
}
