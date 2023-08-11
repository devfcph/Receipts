package storage

import (
	"ReceiptProcessor/models"
	"ReceiptProcessor/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReceiptStorage(t *testing.T) {
	receiptStorage := storage.NewReceiptStorage()

	assert.NotNil(t, receiptStorage)
	assert.Empty(t, receiptStorage.Receipts)
}

func TestReceiptStorage_AddReceipt(t *testing.T) {
	receiptStorage := storage.NewReceiptStorage()

	receipt := &models.Receipt{
		Id:           "123",
		Retailer:     "Example Retailer",
		PurchaseDate: "2023-08-11",
		PurchaseTime: "14:30",
		Items:        nil,
		Total:        "100.00",
	}

	receiptStorage.AddReceipt(receipt)

	assert.Equal(t, 1, len(receiptStorage.Receipts))
	assert.Equal(t, receipt, receiptStorage.Receipts[0])
}

func TestReceiptStorage_GetAllReceipts(t *testing.T) {
	receiptStorage := storage.NewReceiptStorage()

	receipt1 := &models.Receipt{
		Id:           "123",
		Retailer:     "Retailer 1",
		PurchaseDate: "2023-08-11",
		PurchaseTime: "14:30",
		Items:        nil,
		Total:        "100.00",
	}

	receipt2 := &models.Receipt{
		Id:           "456",
		Retailer:     "Retailer 2",
		PurchaseDate: "2023-08-12",
		PurchaseTime: "10:00",
		Items:        nil,
		Total:        "50.00",
	}

	receiptStorage.AddReceipt(receipt1)
	receiptStorage.AddReceipt(receipt2)

	allReceipts := receiptStorage.GetAllReceipts()

	assert.Equal(t, 2, len(allReceipts))
	assert.Equal(t, receipt1, allReceipts[0])
	assert.Equal(t, receipt2, allReceipts[1])
}

func TestReceiptStorage_GetReceiptById(t *testing.T) {
	receiptStorage := storage.NewReceiptStorage()

	receipt1 := &models.Receipt{
		Id:           "123",
		Retailer:     "Retailer 1",
		PurchaseDate: "2023-08-11",
		PurchaseTime: "14:30",
		Items:        nil,
		Total:        "100.00",
	}

	receipt2 := &models.Receipt{
		Id:           "456",
		Retailer:     "Retailer 2",
		PurchaseDate: "2023-08-12",
		PurchaseTime: "10:00",
		Items:        nil,
		Total:        "50.00",
	}

	receiptStorage.AddReceipt(receipt1)
	receiptStorage.AddReceipt(receipt2)

	foundReceipt := receiptStorage.GetReceiptById("456")
	assert.Equal(t, receipt2, foundReceipt)

	notFoundReceipt := receiptStorage.GetReceiptById("789")
	assert.Nil(t, notFoundReceipt)
}
