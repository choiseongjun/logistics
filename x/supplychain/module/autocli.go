package supplychain

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "logistics/api/logistics/supplychain"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ProductAll",
					Use:       "list-product",
					Short:     "List all product",
				},
				{
					RpcMethod:      "Product",
					Use:            "show-product [id]",
					Short:          "Shows a product by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "ShipmentAll",
					Use:       "list-shipment",
					Short:     "List all shipment",
				},
				{
					RpcMethod:      "Shipment",
					Use:            "show-shipment [id]",
					Short:          "Shows a shipment by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "OrderAll",
					Use:       "list-order",
					Short:     "List all order",
				},
				{
					RpcMethod:      "Order",
					Use:            "show-order [id]",
					Short:          "Shows a order by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "TransferOwnership",
					Use:            "transfer-ownership [product-id] [new-owner] [location]",
					Short:          "Send a transfer-ownership tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "productId"}, {ProtoField: "newOwner"}, {ProtoField: "location"}},
				},
				{
					RpcMethod:      "UpdateStatus",
					Use:            "update-status [product-id] [new-status]",
					Short:          "Send a update-status tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "productId"}, {ProtoField: "newStatus"}},
				},
				{
					RpcMethod:      "UpdateShipment",
					Use:            "update-shipment [shipment-id] [new-status] [new-location]",
					Short:          "Send a update-shipment tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "shipmentId"}, {ProtoField: "newStatus"}, {ProtoField: "newLocation"}},
				},
				{
					RpcMethod:      "CreateProduct",
					Use:            "create-product [product-id] [name] [origin] [current-owner] [status]",
					Short:          "Send a create-product tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "productId"}, {ProtoField: "name"}, {ProtoField: "origin"}, {ProtoField: "currentOwner"}, {ProtoField: "status"}},
				},
				{
					RpcMethod:      "CreateOrder",
					Use:            "create-order [orderId] [purchaseId] [seller] [buyer] [status] [items] [price] [timestamp]",
					Short:          "Create order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "orderId"}, {ProtoField: "purchaseId"}, {ProtoField: "seller"}, {ProtoField: "buyer"}, {ProtoField: "status"}, {ProtoField: "items"}, {ProtoField: "price"}, {ProtoField: "timestamp"}},
				},
				{
					RpcMethod:      "UpdateOrder",
					Use:            "update-order [id] [orderId] [purchaseId] [seller] [buyer] [status] [items] [price] [timestamp]",
					Short:          "Update order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "orderId"}, {ProtoField: "purchaseId"}, {ProtoField: "seller"}, {ProtoField: "buyer"}, {ProtoField: "status"}, {ProtoField: "items"}, {ProtoField: "price"}, {ProtoField: "timestamp"}},
				},
				{
					RpcMethod:      "DeleteOrder",
					Use:            "delete-order [id]",
					Short:          "Delete order",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
