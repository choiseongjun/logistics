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

func (k Keeper) ShipmentAll(ctx context.Context, req *types.QueryAllShipmentRequest) (*types.QueryAllShipmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var shipments []types.Shipment

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	shipmentStore := prefix.NewStore(store, types.KeyPrefix(types.ShipmentKey))

	pageRes, err := query.Paginate(shipmentStore, req.Pagination, func(key []byte, value []byte) error {
		var shipment types.Shipment
		if err := k.cdc.Unmarshal(value, &shipment); err != nil {
			return err
		}

		shipments = append(shipments, shipment)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllShipmentResponse{Shipment: shipments, Pagination: pageRes}, nil
}

func (k Keeper) Shipment(ctx context.Context, req *types.QueryGetShipmentRequest) (*types.QueryGetShipmentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	shipment, found := k.GetShipment(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetShipmentResponse{Shipment: shipment}, nil
}
