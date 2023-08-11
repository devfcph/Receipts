package main

import (
	"ReceiptProcessor/config"
	"ReceiptProcessor/controllers"
	"ReceiptProcessor/storage"
	"github.com/gin-gonic/gin"
)

func main() {

	// Create a new instance of the receipt storage.
	receiptStorage := storage.NewReceiptStorage()

	// Create a new instance of the receipt controller and pass in the receipt storage.
	receiptController := controllers.NewReceiptController(receiptStorage)

	// Create a new Gin router with default middleware.
	router := gin.Default()

	// Define a route to handle fetching all receipts.
	router.GET(config.GetDefaultEndPoint()+"/getAll", receiptController.GetReceipts)

	// Define a route to handle processing a receipt.
	router.POST(config.GetDefaultEndPoint()+"/process", receiptController.ProcessReceipt)

	// Define a route to handle fetching points by receipt ID.
	router.GET(config.GetDefaultEndPoint()+"/:id/points", receiptController.GetPointsById)

	// Run the Gin router on the specified server address.
	router.Run(config.GetServer())
}
