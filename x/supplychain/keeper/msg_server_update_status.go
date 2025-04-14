package keeper

import (
	"context"
	"strconv"
	"time"

	errorsmod "cosmossdk.io/errors"

	"logistics/x/supplychain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateStatus(goCtx context.Context, msg *types.MsgUpdateStatus) (*types.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	productId, err := strconv.ParseUint(msg.ProductId, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid product ID format")
	}
	// 제품 존재 여부 확인
	product, found := k.GetProduct(ctx, productId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "product not found")
	}

	// 현재 소유자만 상태 업데이트 가능
	if product.CurrentOwner != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// 상태 이전값 저장
	oldStatus := product.Status

	// 제품 상태 업데이트
	product.Status = msg.NewStatus

	// 제품 정보 저장
	k.SetProduct(ctx, product)

	// 타임스탬프 생성
	timestamp := time.Now().Unix()

	// 이벤트 발생
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeStatusUpdated,
			sdk.NewAttribute(types.AttributeKeyProductId, msg.ProductId),
			sdk.NewAttribute(types.AttributeKeyOldStatus, oldStatus),
			sdk.NewAttribute(types.AttributeKeyNewStatus, msg.NewStatus),
			sdk.NewAttribute(types.AttributeKeyTimestamp, strconv.FormatInt(timestamp, 10)),
		),
	)

	return &types.MsgUpdateStatusResponse{
		ProductId: msg.ProductId,
		Timestamp: strconv.FormatInt(timestamp, 10),
	}, nil
}
