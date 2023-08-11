package models

import "encoding/json"

func UnmarshalTemperatures(data []byte) (Receipt, error) {
	var r Receipt
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Receipt) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Receipt struct {
	Id           string
	Retailer     string        `json:"retailer"`
	PurchaseDate string        `json:"purchaseDate"`
	PurchaseTime string        `json:"purchaseTime"`
	Items        []ReceiptItem `json:"items"`
	Total        string        `json:"total"`
}

func NewReceipt(retailer string, purchaseDate string, purchaseTime string, items []ReceiptItem, total string) *Receipt {
	return &Receipt{Retailer: retailer, PurchaseDate: purchaseDate, PurchaseTime: purchaseTime, Items: items, Total: total}
}
