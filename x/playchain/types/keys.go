package types

const (
	// ModuleName defines the module name
	ModuleName = "playchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_playchain"
)

var (
	ParamsKey = []byte("p_playchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
