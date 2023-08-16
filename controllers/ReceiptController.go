package controllers

import (
	"ReceiptProcessor/models"
	"ReceiptProcessor/storage"
	"ReceiptProcessor/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// ReceiptController handles HTTP requests related to receipts.
type ReceiptController struct {
	receiptStorage *storage.ReceiptStorage
}

// NewReceiptController creates a new instance of ReceiptController.
func NewReceiptController(receiptStorage *storage.ReceiptStorage) *ReceiptController {
	return &ReceiptController{
		receiptStorage: receiptStorage,
	}
}

// GetReceipts handles the GET request to fetch all receipts.
func (controller ReceiptController) GetReceipts(context *gin.Context) {
	receipts := controller.receiptStorage.GetAllReceipts()

	context.IndentedJSON(http.StatusOK, receipts)
}

// GetPointsById handles the GET request to fetch points for a receipt by ID.
func (controller ReceiptController) GetPointsById(context *gin.Context) {
	idParam := context.Param("id")

	receipt := controller.receiptStorage.GetReceiptById(idParam)
	if receipt == nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No receipt found for that id"})
		return
	}

	var newReceiptPoints models.ReceiptPoints
	newReceiptPoints.Points = receipt.Points

	context.IndentedJSON(http.StatusOK, newReceiptPoints)
}

// ProcessReceipt handles the POST request to process a new receipt.
func (controller ReceiptController) ProcessReceipt(context *gin.Context) {
	var newReceipt models.Receipt
	var newReciptGuid models.ReceiptID

	// Bind the JSON request body to the newReceipt struct
	if err := context.ShouldBindJSON(&newReceipt); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The receipt is invalid.  [ " + err.Error() + " ]"})
		return
	}

	newId, _ := uuid.NewUUID()
	newReceipt.Id = newId.String()
	newReciptGuid.ID = newReceipt.Id

	newReceipt.Points = utils.CalculateReceiptPoints(&newReceipt)

	controller.receiptStorage.AddReceipt(&newReceipt)

	context.IndentedJSON(http.StatusOK, newReciptGuid)
}
