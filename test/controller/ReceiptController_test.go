package controller

import (
	"ReceiptProcessor/controllers"
	"ReceiptProcessor/models"
	"ReceiptProcessor/storage"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReceiptController_GetReceipts(t *testing.T) {
	receiptStorage := storage.NewReceiptStorage()
	controller := controllers.NewReceiptController(receiptStorage)

	router := gin.Default()
	router.GET("/receipts", controller.GetReceipts)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/receipts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestReceiptController_GetPointsById(t *testing.T) {
	receiptStorage := storage.NewReceiptStorage()
	controller := controllers.NewReceiptController(receiptStorage)

	receipt := &models.Receipt{
		Id:           "123",
		Retailer:     "Retailer 1",
		PurchaseDate: "2023-08-11",
		PurchaseTime: "14:30",
		Items:        nil,
		Total:        "100.00",
	}

	receiptStorage.AddReceipt(receipt)

	router := gin.Default()
	router.GET("/receipts/:id/points", controller.GetPointsById)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/receipts/123/points", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestReceiptController_ProcessReceipt(t *testing.T) {
	receiptStorage := storage.NewReceiptStorage()
	controller := controllers.NewReceiptController(receiptStorage)

	router := gin.Default()
	router.POST("/receipts", controller.ProcessReceipt)

	receiptJSON := `{"retailer": "Retailer 1", "purchaseDate": "2023-08-11", "purchaseTime": "14:30", "total": "100.00", "items": []}`
	bodyReader := strings.NewReader(receiptJSON)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/receipts", bodyReader)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}
