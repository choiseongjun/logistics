package types

const (
	// ModuleName defines the module name
	ModuleName = "supplychain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_supplychain"
)

var (
	ParamsKey = []byte("p_supplychain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ProductKey      = "Product/value/"
	ProductCountKey = "Product/count/"
)

const (
	ShipmentKey      = "Shipment/value/"
	ShipmentCountKey = "Shipment/count/"
)

const (
	OrderKey      = "Order/value/"
	OrderCountKey = "Order/count/"
)
