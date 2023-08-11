package utils

import (
	"ReceiptProcessor/constants"
	"ReceiptProcessor/models"
	"ReceiptProcessor/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateReceiptPoints(t *testing.T) {
	receipt := &models.Receipt{
		Id:           "123",
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.ReceiptItem{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			}, {
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			}, {
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			}, {
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			}, {
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			},
		},
		Total: "35.35",
	}

	expectedPoints := int64(28)
	calculatedPoints := utils.CalculateReceiptPoints(receipt)

	assert.Equal(t, expectedPoints, calculatedPoints)
}

func TestGetPointsByRetailerName(t *testing.T) {
	retailerName := "Target"

	points := utils.GetPointsByRetailerName(retailerName)

	assert.Equal(t, int64(len(retailerName)), points)
}

func TestGetPointsByTotal(t *testing.T) {
	total := "100.00"

	points := utils.GetPointsByTotal(total)

	assert.Equal(t, int64(constants.FIFTY+constants.TWENTYFIVE), points)
}

func TestGetPointsByReceiptItems(t *testing.T) {
	receipt := &models.Receipt{
		Id:           "123",
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.ReceiptItem{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			}, {
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			}, {
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			}, {
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			}, {
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			},
		},
		Total: "35.35",
	}

	expectedPoints := int64(16)
	calculatedPoints := utils.GetPointsByReceiptItems(receipt)

	assert.Equal(t, expectedPoints, calculatedPoints)
}

func TestGetPointsByPurchaseInfo(t *testing.T) {
	date := "2023-08-11"
	time := "14:30"

	expectedPoints := int64(constants.SIX + constants.TEN)
	calculatedPoints := utils.GetPointsByPurchaseInfo(date, time)

	assert.Equal(t, expectedPoints, calculatedPoints)
}

func TestGetPointsByItemDescription(t *testing.T) {
	item := models.ReceiptItem{
		ShortDescription: "Item 1",
		Price:            "10.00",
	}

	expectedPoints := int64(2)
	calculatedPoints := utils.GetPointsByItemDescription(item)

	assert.Equal(t, expectedPoints, calculatedPoints)
}

func TestGetPointsByNumberOfItems(t *testing.T) {
	totalItems := 4

	expectedPoints := int64(10)
	calculatedPoints := utils.GetPointsByNumberOfItems(totalItems)

	assert.Equal(t, expectedPoints, calculatedPoints)
}

func TestIsMultipleOf(t *testing.T) {
	assert.True(t, utils.IsMultipleOf(6, 3))
	assert.False(t, utils.IsMultipleOf(7, 3))
}

func TestIsDateValue(t *testing.T) {
	validDate := "2023-08-11"
	invalidDate := "2023-08-50"

	assert.True(t, utils.IsDateValue(validDate))
	assert.False(t, utils.IsDateValue(invalidDate))
}

func TestIsHourValue(t *testing.T) {
	validTime := "14:30"
	invalidTime := "25:00"

	assert.True(t, utils.IsHourValue(validTime))
	assert.False(t, utils.IsHourValue(invalidTime))
}

func TestGetPointsByDayFromDateString(t *testing.T) {
	date := "2023-08-11"

	expectedPoints := int64(constants.SIX)
	calculatedPoints := utils.GetPointsByDayFromDateString(date)

	assert.Equal(t, expectedPoints, calculatedPoints)
}

func TestGetPointsByPurchaseTime(t *testing.T) {
	time := "14:30"

	expectedPoints := int64(constants.TEN)
	calculatedPoints := utils.GetPointsByPurchaseTime(time)

	assert.Equal(t, expectedPoints, calculatedPoints)
}
