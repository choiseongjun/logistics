package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"logistics/x/supplychain/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateOrder(goCtx context.Context, msg *types.MsgCreateOrder) (*types.MsgCreateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 제품 ID를 uint64로 변환
	productId, err := strconv.ParseUint(msg.Items, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid product ID format")
	}

	// 제품 존재 여부 확인
	product, found := k.GetProduct(ctx, productId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "product not found")
	}

	// 판매자가 제품 소유자인지 확인
	if product.CurrentOwner != msg.Seller {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "seller is not the product owner")
	}

	// 현재 시간을 타임스탬프로 설정
	timestamp := time.Now().Format(time.RFC3339)

	var order = types.Order{
		Creator:    msg.Creator,
		OrderId:    msg.OrderId,
		PurchaseId: msg.PurchaseId,
		Seller:     msg.Seller,
		Buyer:      msg.Buyer,
		Status:     "pending", // 초기 상태를 pending으로 설정
		Items:      msg.Items, // 제품 ID
		Price:      msg.Price,
		Timestamp:  timestamp,
	}

	// 주문 저장
	id := k.AppendOrder(ctx, order)

	// 이벤트 발생
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeOrderCreated,
			sdk.NewAttribute(types.AttributeKeyOrderId, msg.OrderId),
			sdk.NewAttribute(types.AttributeKeyProductId, msg.Items),
			sdk.NewAttribute(types.AttributeKeySeller, msg.Seller),
			sdk.NewAttribute(types.AttributeKeyBuyer, msg.Buyer),
			sdk.NewAttribute(types.AttributeKeyPrice, msg.Price),
			sdk.NewAttribute(types.AttributeKeyTimestamp, timestamp),
		),
	)

	return &types.MsgCreateOrderResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateOrder(goCtx context.Context, msg *types.MsgUpdateOrder) (*types.MsgUpdateOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var order = types.Order{
		Creator:    msg.Creator,
		Id:         msg.Id,
		OrderId:    msg.OrderId,
		PurchaseId: msg.PurchaseId,
		Seller:     msg.Seller,
		Buyer:      msg.Buyer,
		Status:     msg.Status,
		Items:      msg.Items,
		Price:      msg.Price,
		Timestamp:  msg.Timestamp,
	}

	// Checks that the element exists
	val, found := k.GetOrder(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetOrder(ctx, order)

	return &types.MsgUpdateOrderResponse{}, nil
}

func (k msgServer) DeleteOrder(goCtx context.Context, msg *types.MsgDeleteOrder) (*types.MsgDeleteOrderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetOrder(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveOrder(ctx, msg.Id)

	return &types.MsgDeleteOrderResponse{}, nil
}
