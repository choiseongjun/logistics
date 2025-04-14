package supplychain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"logistics/x/supplychain/keeper"
	"logistics/x/supplychain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the product
	for _, elem := range genState.ProductList {
		k.SetProduct(ctx, elem)
	}

	// Set product count
	k.SetProductCount(ctx, genState.ProductCount)
	// Set all the shipment
	for _, elem := range genState.ShipmentList {
		k.SetShipment(ctx, elem)
	}

	// Set shipment count
	k.SetShipmentCount(ctx, genState.ShipmentCount)
	// Set all the order
	for _, elem := range genState.OrderList {
		k.SetOrder(ctx, elem)
	}

	// Set order count
	k.SetOrderCount(ctx, genState.OrderCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ProductList = k.GetAllProduct(ctx)
	genesis.ProductCount = k.GetProductCount(ctx)
	genesis.ShipmentList = k.GetAllShipment(ctx)
	genesis.ShipmentCount = k.GetShipmentCount(ctx)
	genesis.OrderList = k.GetAllOrder(ctx)
	genesis.OrderCount = k.GetOrderCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
