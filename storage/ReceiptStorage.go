package storage

import (
	"ReceiptProcessor/models"
	"sync"
)

var data []models.Receipt

type ReceiptStorage struct {
	receipts []*models.Receipt
	mutex    sync.Mutex
}

func NewReceiptStorage() *ReceiptStorage {
	return &ReceiptStorage{
		receipts: make([]*models.Receipt, 0),
	}
}

func (storage *ReceiptStorage) AddReceipt(receipt *models.Receipt) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	storage.receipts = append(storage.receipts, receipt)
}

func (storage *ReceiptStorage) GetAllReceipts() []*models.Receipt {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	return storage.receipts
}

func (storage *ReceiptStorage) GetReceiptById(id string) *models.Receipt {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	for _, receipt := range storage.receipts {
		if receipt.Id == id {
			return receipt
		}
	}

	return nil
}
