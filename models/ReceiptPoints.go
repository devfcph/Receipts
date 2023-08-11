package models

type ReceiptPoints struct {
	Points int64 `json:"points"`
}

func NewReceiptPoints(points int64) *ReceiptPoints {
	return &ReceiptPoints{Points: points}
}
