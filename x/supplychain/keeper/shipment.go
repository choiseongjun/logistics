package keeper

import (
	"context"
	"encoding/binary"

	"logistics/x/supplychain/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetShipmentCount get the total number of shipment
func (k Keeper) GetShipmentCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ShipmentCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetShipmentCount set the total number of shipment
func (k Keeper) SetShipmentCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ShipmentCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendShipment appends a shipment in the store with a new id and update the count
func (k Keeper) AppendShipment(
	ctx context.Context,
	shipment types.Shipment,
) uint64 {
	// Create the shipment
	count := k.GetShipmentCount(ctx)

	// Set the ID of the appended value
	shipment.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShipmentKey))
	appendedValue := k.cdc.MustMarshal(&shipment)
	store.Set(GetShipmentIDBytes(shipment.Id), appendedValue)

	// Update shipment count
	k.SetShipmentCount(ctx, count+1)

	return count
}

// SetShipment set a specific shipment in the store
func (k Keeper) SetShipment(ctx context.Context, shipment types.Shipment) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShipmentKey))
	b := k.cdc.MustMarshal(&shipment)
	store.Set(GetShipmentIDBytes(shipment.Id), b)
}

// GetShipment returns a shipment from its id
func (k Keeper) GetShipment(ctx context.Context, id uint64) (val types.Shipment, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShipmentKey))
	b := store.Get(GetShipmentIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveShipment removes a shipment from the store
func (k Keeper) RemoveShipment(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShipmentKey))
	store.Delete(GetShipmentIDBytes(id))
}

// GetAllShipment returns all shipment
func (k Keeper) GetAllShipment(ctx context.Context) (list []types.Shipment) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ShipmentKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Shipment
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetShipmentIDBytes returns the byte representation of the ID
func GetShipmentIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.ShipmentKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
