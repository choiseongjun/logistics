package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ProductList:  []Product{},
		ShipmentList: []Shipment{},
		OrderList:    []Order{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in product
	productIdMap := make(map[uint64]bool)
	productCount := gs.GetProductCount()
	for _, elem := range gs.ProductList {
		if _, ok := productIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for product")
		}
		if elem.Id >= productCount {
			return fmt.Errorf("product id should be lower or equal than the last id")
		}
		productIdMap[elem.Id] = true
	}
	// Check for duplicated ID in shipment
	shipmentIdMap := make(map[uint64]bool)
	shipmentCount := gs.GetShipmentCount()
	for _, elem := range gs.ShipmentList {
		if _, ok := shipmentIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for shipment")
		}
		if elem.Id >= shipmentCount {
			return fmt.Errorf("shipment id should be lower or equal than the last id")
		}
		shipmentIdMap[elem.Id] = true
	}
	// Check for duplicated ID in order
	orderIdMap := make(map[uint64]bool)
	orderCount := gs.GetOrderCount()
	for _, elem := range gs.OrderList {
		if _, ok := orderIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for order")
		}
		if elem.Id >= orderCount {
			return fmt.Errorf("order id should be lower or equal than the last id")
		}
		orderIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
