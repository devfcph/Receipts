package utils

import (
	"ReceiptProcessor/models"
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func CalculateReceiptPoints(receipt *models.Receipt) int64 {
	totalPoints := int64(0)
	totalPoints = getPointsByRetailerName(receipt.Retailer) + getPointsByTotal(receipt.Total)
	return totalPoints
}

func getPointsByRetailerName(retailName string) int64 {
	points := int64(0)

	for _, char := range retailName {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}
	return points
}

func getPointsByTotal(total string) int64 {

	points := int64(0)

	if totalParsed, err := strconv.ParseFloat(total, 64); err == nil {
		fmt.Println(totalParsed)
		if math.Ceil(totalParsed) == totalParsed {
			points = int64(50)
		}

		if isMultipleOf(totalParsed, 0.25) {
			points = points + int64(25)
		}
	}

	return points
}

func isMultipleOf(number float64, multipleOf float64) bool {
	if multipleOf == 0 {
		return false
	}
	quotient := number / multipleOf
	return math.Mod(quotient, 1) == 0
}
