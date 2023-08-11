package models

// ReceiptPoints represents the points earned from a receipt.
type ReceiptPoints struct {
	Points int64 `json:"points"` // The number of points earned.
}

// NewReceiptPoints creates a new instance of ReceiptPoints with the provided points value.
func NewReceiptPoints(points int64) *ReceiptPoints {
	return &ReceiptPoints{Points: points}
}
