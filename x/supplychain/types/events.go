// types/events.go - 이벤트 타입 및 속성 키 정의
package types

// 이벤트 타입 상수
const (
	EventTypeOwnershipTransferred = "ownership_transferred"
	EventTypeProductCreated       = "product_created"
	EventTypeNewOrder             = "new_order"
	EventTypeShipmentUpdated      = "shipment_updated"
	// 다른 이벤트 타입 추가 가능
	EventTypeOrderCreated = "order_created"
	AttributeKeyOrderId   = "order_id"
	AttributeKeySeller    = "seller"
	AttributeKeyBuyer     = "buyer"
	AttributeKeyPrice     = "price"
)

// 이벤트 속성 키 상수
const (
	AttributeKeyProductId     = "product_id"
	AttributeKeyPreviousOwner = "previous_owner"
	AttributeKeyNewOwner      = "new_owner"
	AttributeKeyTimestamp     = "timestamp"
	AttributeKeyLocation      = "location"
	AttributeKeyShipmentId    = "shipment_id"
	AttributeKeyOldStatus     = "old_status"
	AttributeKeyNewStatus     = "new_status"
	EventTypeStatusUpdated    = "status_updated"
	AttributeKeyName          = "name"
	AttributeKeyOrigin        = "origin"
	AttributeKeyOwner         = "owner"
	AttributeKeyStatus        = "key_status"
	// 다른 속성 키 추가 가능
)
