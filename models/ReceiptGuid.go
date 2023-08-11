package models

// ReceiptID represents a unique identifier for a receipt.
type ReceiptID struct {
	ID string `json:"id"` // The ID value.
}

// NewReceiptID creates a new instance of ReceiptID with the provided ID value.
func NewReceiptID(ID string) *ReceiptID {
	return &ReceiptID{ID: ID}
}
