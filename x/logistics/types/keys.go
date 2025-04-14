package types

const (
	// ModuleName defines the module name
	ModuleName = "logistics"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_logistics"
)

var (
	ParamsKey = []byte("p_logistics")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
