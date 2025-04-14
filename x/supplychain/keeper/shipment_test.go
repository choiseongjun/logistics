package keeper_test

import (
	"context"
	"testing"

	keepertest "logistics/testutil/keeper"
	"logistics/testutil/nullify"
	"logistics/x/supplychain/keeper"
	"logistics/x/supplychain/types"

	"github.com/stretchr/testify/require"
)

func createNShipment(keeper keeper.Keeper, ctx context.Context, n int) []types.Shipment {
	items := make([]types.Shipment, n)
	for i := range items {
		items[i].Id = keeper.AppendShipment(ctx, items[i])
	}
	return items
}

func TestShipmentGet(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	items := createNShipment(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetShipment(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestShipmentRemove(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	items := createNShipment(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveShipment(ctx, item.Id)
		_, found := keeper.GetShipment(ctx, item.Id)
		require.False(t, found)
	}
}

func TestShipmentGetAll(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	items := createNShipment(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllShipment(ctx)),
	)
}

func TestShipmentCount(t *testing.T) {
	keeper, ctx := keepertest.SupplychainKeeper(t)
	items := createNShipment(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetShipmentCount(ctx))
}
