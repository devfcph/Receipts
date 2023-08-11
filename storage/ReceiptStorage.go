package storage

import (
	"ReceiptProcessor/models"
	"sync"
)

// ReceiptStorage is responsible for managing receipts in memory.
type ReceiptStorage struct {
	Receipts []*models.Receipt // A slice to hold receipt instances.
	mutex    sync.Mutex        // Mutex to manage concurrent access.
}

// NewReceiptStorage creates a new instance of ReceiptStorage.
func NewReceiptStorage() *ReceiptStorage {
	return &ReceiptStorage{
		Receipts: make([]*models.Receipt, 0),
	}
}

// AddReceipt adds a new receipt to the storage.
func (storage *ReceiptStorage) AddReceipt(receipt *models.Receipt) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	storage.Receipts = append(storage.Receipts, receipt)
}

// GetAllReceipts returns all receipts stored.
func (storage *ReceiptStorage) GetAllReceipts() []*models.Receipt {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	return storage.Receipts
}

// GetReceiptById retrieves a receipt by its ID.
func (storage *ReceiptStorage) GetReceiptById(id string) *models.Receipt {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	for _, receipt := range storage.Receipts {
		if receipt.Id == id {
			return receipt
		}
	}

	return nil
}
