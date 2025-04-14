package supplychain_test

import (
	"testing"

	keepertest "logistics/testutil/keeper"
	"logistics/testutil/nullify"
	supplychain "logistics/x/supplychain/module"
	"logistics/x/supplychain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ProductList: []types.Product{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ProductCount: 2,
		ShipmentList: []types.Shipment{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ShipmentCount: 2,
		OrderList: []types.Order{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		OrderCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SupplychainKeeper(t)
	supplychain.InitGenesis(ctx, k, genesisState)
	got := supplychain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProductList, got.ProductList)
	require.Equal(t, genesisState.ProductCount, got.ProductCount)
	require.ElementsMatch(t, genesisState.ShipmentList, got.ShipmentList)
	require.Equal(t, genesisState.ShipmentCount, got.ShipmentCount)
	require.ElementsMatch(t, genesisState.OrderList, got.OrderList)
	require.Equal(t, genesisState.OrderCount, got.OrderCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
