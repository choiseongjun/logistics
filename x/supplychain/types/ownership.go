package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// 소유권 변경 기록
type OwnershipChange struct {
	ProductId     string         `json:"product_id"`
	PreviousOwner sdk.AccAddress `json:"previous_owner"`
	NewOwner      sdk.AccAddress `json:"new_owner"`
	Timestamp     time.Time      `json:"timestamp"`
	Location      string         `json:"location,omitempty"`
	TransactionId string         `json:"transaction_id"`
}
