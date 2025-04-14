package keeper

import (
	"context"

	"logistics/x/supplychain/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) OrderAll(ctx context.Context, req *types.QueryAllOrderRequest) (*types.QueryAllOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var orders []types.Order

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	orderStore := prefix.NewStore(store, types.KeyPrefix(types.OrderKey))

	pageRes, err := query.Paginate(orderStore, req.Pagination, func(key []byte, value []byte) error {
		var order types.Order
		if err := k.cdc.Unmarshal(value, &order); err != nil {
			return err
		}

		orders = append(orders, order)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllOrderResponse{Order: orders, Pagination: pageRes}, nil
}

func (k Keeper) Order(ctx context.Context, req *types.QueryGetOrderRequest) (*types.QueryGetOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	order, found := k.GetOrder(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetOrderResponse{Order: order}, nil
}
