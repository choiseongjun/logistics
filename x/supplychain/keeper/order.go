package keeper

import (
	"context"
	"encoding/binary"

	"logistics/x/supplychain/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetOrderCount get the total number of order
func (k Keeper) GetOrderCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.OrderCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetOrderCount set the total number of order
func (k Keeper) SetOrderCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.OrderCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendOrder appends a order in the store with a new id and update the count
func (k Keeper) AppendOrder(
	ctx context.Context,
	order types.Order,
) uint64 {
	// Create the order
	count := k.GetOrderCount(ctx)

	// Set the ID of the appended value
	order.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrderKey))
	appendedValue := k.cdc.MustMarshal(&order)
	store.Set(GetOrderIDBytes(order.Id), appendedValue)

	// Update order count
	k.SetOrderCount(ctx, count+1)

	return count
}

// SetOrder set a specific order in the store
func (k Keeper) SetOrder(ctx context.Context, order types.Order) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrderKey))
	b := k.cdc.MustMarshal(&order)
	store.Set(GetOrderIDBytes(order.Id), b)
}

// GetOrder returns a order from its id
func (k Keeper) GetOrder(ctx context.Context, id uint64) (val types.Order, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrderKey))
	b := store.Get(GetOrderIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrder removes a order from the store
func (k Keeper) RemoveOrder(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrderKey))
	store.Delete(GetOrderIDBytes(id))
}

// GetAllOrder returns all order
func (k Keeper) GetAllOrder(ctx context.Context) (list []types.Order) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OrderKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Order
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOrderIDBytes returns the byte representation of the ID
func GetOrderIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.OrderKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
