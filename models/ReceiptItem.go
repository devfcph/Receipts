package models

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func NewReceiptItem(shortDescription string, price string) *ReceiptItem {
	return &ReceiptItem{ShortDescription: shortDescription, Price: price}
}
