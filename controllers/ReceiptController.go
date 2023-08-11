package controllers

import (
	"ReceiptProcessor/models"
	"ReceiptProcessor/storage"
	"ReceiptProcessor/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type ReceiptController struct {
	receiptStorage *storage.ReceiptStorage
}

func NewReceiptController(receiptStorage *storage.ReceiptStorage) *ReceiptController {
	return &ReceiptController{
		receiptStorage: receiptStorage,
	}
}

func (controller ReceiptController) GetReceipts(context *gin.Context) {
	receipts := controller.receiptStorage.GetAllReceipts()

	context.IndentedJSON(http.StatusOK, receipts)
}

func (controller ReceiptController) GetPointsById(context *gin.Context) {
	idParam := context.Param("id")

	receipt := controller.receiptStorage.GetReceiptById(idParam)
	if receipt == nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No receipt found for that id"})
		return
	}

	var newReceiptPoints models.ReceiptPoints
	newReceiptPoints.Points = utils.CalculateReceiptPoints(receipt)

	context.IndentedJSON(http.StatusOK, newReceiptPoints)
}

func (controller ReceiptController) ProcessReceipt(context *gin.Context) {
	var newReceipt models.Receipt
	var newReciptGuid models.ReceiptID

	if err := context.ShouldBindJSON(&newReceipt); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "The receipt is invalid.  [ " + err.Error() + " ]"})
		return
	}

	newId, _ := uuid.NewUUID()
	newReceipt.Id = newId.String()
	newReciptGuid.ID = newReceipt.Id

	controller.receiptStorage.AddReceipt(&newReceipt)

	context.IndentedJSON(http.StatusOK, newReciptGuid)
}
