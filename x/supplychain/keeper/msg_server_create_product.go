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

func (k msgServer) CreateProduct(goCtx context.Context, msg *types.MsgCreateProduct) (*types.MsgCreateProductResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	productId, err := strconv.ParseUint(msg.ProductId, 10, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Invalid product ID format, must be a numeric string")
	}

	// 제품 ID가 이미 존재하는지 확인
	_, found := k.GetProduct(ctx, productId)
	if found {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Product with this ID already exists")
	}
	// 제품 ID가 이미 존재하는지 확인
	//_, found := k.GetProduct(ctx, productId)
	//if found {
	//	return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Product with this ID already exists")
	//}

	// 현재 시간을 타임스탬프로 사용
	timestamp := time.Now().Unix()

	// 새 제품 객체 생성
	product := types.Product{
		ProductId:    msg.ProductId,
		Name:         msg.Name,
		Origin:       msg.Origin,
		CurrentOwner: msg.CurrentOwner,
		Status:       msg.Status,
		Creator:      msg.Creator,
	}

	// 스토리지에 제품 저장
	k.SetProduct(ctx, product)

	// 이벤트 생성 및 발생
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeProductCreated,
			sdk.NewAttribute(types.AttributeKeyProductId, msg.ProductId),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyOrigin, msg.Origin),
			sdk.NewAttribute(types.AttributeKeyOwner, msg.CurrentOwner),
			sdk.NewAttribute(types.AttributeKeyStatus, msg.Status),
			sdk.NewAttribute(types.AttributeKeyTimestamp, string(timestamp)),
		),
	)

	return &types.MsgCreateProductResponse{
		ProductId: msg.ProductId,
	}, nil
}
