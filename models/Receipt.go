package models

// Receipt represents a receipt containing purchase information.
type Receipt struct {
	Id           string        // Unique identifier of the receipt.
	Retailer     string        `json:"retailer"`     // Name of the retailer.
	PurchaseDate string        `json:"purchaseDate"` // Date of the purchase.
	PurchaseTime string        `json:"purchaseTime"` // Time of the purchase.
	Items        []ReceiptItem `json:"items"`        // List of items on the receipt.
	Total        string        `json:"total"`        // Total amount of the purchase.
	Points       int64         // Points
}

// NewReceipt creates a new instance of Receipt.
func NewReceipt(retailer string, purchaseDate string, purchaseTime string, items []ReceiptItem, total string) *Receipt {
	return &Receipt{Retailer: retailer, PurchaseDate: purchaseDate, PurchaseTime: purchaseTime, Items: items, Total: total}
}
