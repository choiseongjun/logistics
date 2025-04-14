package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "logistics/testutil/keeper"
	"logistics/testutil/nullify"
	"logistics/x/supplychain/types"
)

func TestShipmentQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	msgs := createNShipment(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetShipmentRequest
		response *types.QueryGetShipmentResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetShipmentRequest{Id: msgs[0].Id},
			response: &types.QueryGetShipmentResponse{Shipment: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetShipmentRequest{Id: msgs[1].Id},
			response: &types.QueryGetShipmentResponse{Shipment: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetShipmentRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Shipment(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestShipmentQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	msgs := createNShipment(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllShipmentRequest {
		return &types.QueryAllShipmentRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ShipmentAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Shipment), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Shipment),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ShipmentAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Shipment), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Shipment),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ShipmentAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Shipment),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ShipmentAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
