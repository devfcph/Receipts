package models

type ReceiptID struct {
	ID string `json:"id"`
}

func NewReceiptID(ID string) *ReceiptID {
	return &ReceiptID{ID: ID}
}
