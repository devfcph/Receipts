package models

// ReceiptItem represents an item on a receipt.
type ReceiptItem struct {
	ShortDescription string `json:"shortDescription"` // A short description of the item.
	Price            string `json:"price"`            // The price of the item.
}

// NewReceiptItem creates a new instance of ReceiptItem with the provided attributes.
func NewReceiptItem(shortDescription string, price string) *ReceiptItem {
	return &ReceiptItem{ShortDescription: shortDescription, Price: price}
}
