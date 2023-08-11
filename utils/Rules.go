package utils

import (
	"ReceiptProcessor/constants"
	"ReceiptProcessor/models"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// CalculateReceiptPoints calculates the total points earned from a receipt.
func CalculateReceiptPoints(receipt *models.Receipt) int64 {
	totalPoints := int64(0)
	totalPoints += getPointsByRetailerName(receipt.Retailer)
	totalPoints += getPointsByTotal(receipt.Total)
	totalPoints += getPointsByReceiptItems(receipt)
	totalPoints += getPointsByPurchaseInfo(receipt.PurchaseDate, receipt.PurchaseTime)
	return totalPoints
}

// getPointsByRetailerName calculates points based on retailer name.
func getPointsByRetailerName(retailName string) int64 {
	points := int64(0)

	for _, char := range retailName {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}
	return points
}

// getPointsByTotal calculates points based on total amount.
func getPointsByTotal(total string) int64 {

	points := int64(0)

	totalParsed, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return points
	}

	if math.Ceil(totalParsed) == totalParsed {
		points += int64(constants.FIFTY)
	}

	if isMultipleOf(totalParsed, constants.ONEQUARTER) {
		points += int64(constants.TWENTYFIVE)
	}

	return points
}

// getPointsByReceiptItems calculates points based on receipt items.
func getPointsByReceiptItems(receipt *models.Receipt) int64 {
	points := int64(0)

	points += getPointsByNumberOfItems(len(receipt.Items))

	for _, item := range receipt.Items {
		points += getPointsByItemDescription(item)
	}

	return points
}

// getPointsByPurchaseInfo calculates points based on purchase date and time.
func getPointsByPurchaseInfo(date string, time string) int64 {
	points := int64(0)
	if isDateValue(date) {
		points += getPointsByDayFromDateString(date)
	}

	if isHourValue(time) {
		points += getPointsByPurchaseTime(time)
	}

	return points
}

// getPointsByItemDescription calculates points based on item description.
func getPointsByItemDescription(item models.ReceiptItem) int64 {
	points := int64(0)
	factor := constants.ONEFIFTH

	if isMultipleOf(float64(len(strings.TrimSpace(item.ShortDescription))), float64(constants.THREE)) {

		price, err := strconv.ParseFloat(item.Price, 64)
		if err == nil {
			aux := factor * price
			decimals := aux - math.Round(aux)
			_points := int64(math.Round(aux))
			if decimals < 0.5 {
				points = _points + 1
			}
		}

	}

	return points
}

// getPointsByNumberOfItems calculates points based on the number of receipt items.
func getPointsByNumberOfItems(totalItems int) int64 {
	return int64((totalItems / 2) * constants.FIVE)
}

// isMultipleOf checks if a number is a multiple of another number.
func isMultipleOf(number float64, multipleOf float64) bool {
	if multipleOf == 0 {
		return false
	}
	quotient := number / multipleOf
	return math.Mod(quotient, 1) == 0
}

// isDateValue checks if a string represents a valid date.
func isDateValue(stringDate string) bool {
	_, err := time.Parse("2006-01-02", stringDate)
	return err == nil
}

// isHourValue checks if a string represents a valid hour.
func isHourValue(timeString string) bool {

	purchaseTime := strings.Split(timeString, ":")
	if len(purchaseTime) != 2 {
		return false
	}

	hour, err := strconv.Atoi(purchaseTime[0])

	if err != nil {
		return false
	}

	minute, err := strconv.Atoi(purchaseTime[1])

	if err != nil {
		return false
	}

	return (hour >= 0 && hour < 24) && (minute >= 0 && minute < 60)
}

// getPointsByDayFromDateString calculates points based on the day of the purchase date.
func getPointsByDayFromDateString(date string) int64 {
	points := int64(0)

	dateParsed, err := time.Parse("2006-01-02", date)

	if err == nil && dateParsed.Day()%2 == 1 {
		points = int64(constants.SIX)
	}

	return points
}

// getPointsByPurchaseTime calculates points based on the purchase time.
func getPointsByPurchaseTime(purchaseTime string) int64 {
	points := int64(0)

	timeSplit := strings.Split(purchaseTime, ":")

	hour, err := strconv.Atoi(timeSplit[0])
	if err != nil {
		return points
	}

	minute, err := strconv.Atoi(timeSplit[1])
	if err != nil {
		return points
	}

	if (hour >= 14 && hour <= 15) && (minute >= 0 && minute < 60) {
		points = int64(constants.TEN)
	}

	return points

}
