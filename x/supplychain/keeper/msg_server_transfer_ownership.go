package keeper

import (
	"context"
	"strconv"
	"time"

	"logistics/x/supplychain/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) TransferOwnership(goCtx context.Context, msg *types.MsgTransferOwnership) (*types.MsgTransferOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Convert string ProductId to uint64
	productId, err := strconv.ParseUint(msg.ProductId, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid product ID format")
	}

	// 제품 존재 여부 확인
	product, found := k.GetProduct(ctx, productId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "product not found")
	}

	// 현재 소유자만 소유권 이전 가능
	if product.CurrentOwner != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// 소유권 이력에 추가할 데이터 생성
	//unixTimestamp := time.Now().Unix()
	currentTime := time.Now()
	unixTimestamp := currentTime.Unix()

	ownershipChange := types.OwnershipChange{
		PreviousOwner: sdk.AccAddress(product.CurrentOwner),
		NewOwner:      sdk.AccAddress(msg.NewOwner),
		Timestamp:     currentTime, // Unix 타임스탬프를 time.Time으로 변환
		Location:      msg.Location,
	}

	// 제품 정보 업데이트
	product.CurrentOwner = msg.NewOwner
	product.Status = "ownership_transferred"

	// 소유권 이력 업데이트 (필요시 types.proto에 OwnershipChange 타입 및 History 필드 추가 필요)
	// product.History = append(product.History, ownershipChange)

	// 제품 정보 저장
	k.SetProduct(ctx, product)

	// 이벤트 발생
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeOwnershipTransferred,
			sdk.NewAttribute(types.AttributeKeyProductId, msg.ProductId),
			sdk.NewAttribute(types.AttributeKeyPreviousOwner, string(ownershipChange.PreviousOwner)),
			sdk.NewAttribute(types.AttributeKeyNewOwner, string(ownershipChange.NewOwner)),
			sdk.NewAttribute(types.AttributeKeyTimestamp, strconv.FormatInt(unixTimestamp, 10)),
			sdk.NewAttribute(types.AttributeKeyLocation, ownershipChange.Location),
		),
	)

	return &types.MsgTransferOwnershipResponse{
		OwnerId:   msg.NewOwner,
		Timestamp: strconv.FormatInt(unixTimestamp, 10),
	}, nil
}
