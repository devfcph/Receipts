package main

import (
	"ReceiptProcessor/config"
	"ReceiptProcessor/controllers"
	"ReceiptProcessor/storage"
	"github.com/gin-gonic/gin"
)

func main() {

	receiptStorage := storage.NewReceiptStorage()
	receiptController := controllers.NewReceiptController(receiptStorage)

	router := gin.Default()
	router.GET(config.GetDefaultEndPoint()+"/getAll", receiptController.GetReceipts)

	// Process Receipt [POST]
	router.POST(config.GetDefaultEndPoint()+"/process", receiptController.ProcessReceipt)

	//Get points
	router.GET(config.GetDefaultEndPoint()+"/:id/points", receiptController.GetPointsById)

	router.Run(config.GetServer())
}
