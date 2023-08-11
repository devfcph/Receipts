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
	totalPoints += GetPointsByRetailerName(receipt.Retailer)
	totalPoints += GetPointsByTotal(receipt.Total)
	totalPoints += GetPointsByReceiptItems(receipt)
	totalPoints += GetPointsByPurchaseInfo(receipt.PurchaseDate, receipt.PurchaseTime)
	return totalPoints
}

// GetPointsByRetailerName calculates points based on retailer name.
func GetPointsByRetailerName(retailName string) int64 {
	points := int64(0)

	for _, char := range retailName {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}
	return points
}

// GetPointsByTotal calculates points based on total amount.
func GetPointsByTotal(total string) int64 {

	points := int64(0)

	totalParsed, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return points
	}

	if math.Ceil(totalParsed) == totalParsed {
		points += int64(constants.FIFTY)
	}

	if IsMultipleOf(totalParsed, constants.ONEQUARTER) {
		points += int64(constants.TWENTYFIVE)
	}

	return points
}

// GetPointsByReceiptItems calculates points based on receipt items.
func GetPointsByReceiptItems(receipt *models.Receipt) int64 {
	points := int64(0)

	points += GetPointsByNumberOfItems(len(receipt.Items))

	for _, item := range receipt.Items {
		points += GetPointsByItemDescription(item)
	}

	return points
}

// GetPointsByPurchaseInfo calculates points based on purchase date and time.
func GetPointsByPurchaseInfo(date string, time string) int64 {
	points := int64(0)
	if IsDateValue(date) {
		points += GetPointsByDayFromDateString(date)
	}

	if IsHourValue(time) {
		points += GetPointsByPurchaseTime(time)
	}

	return points
}

// GetPointsByItemDescription calculates points based on item description.
func GetPointsByItemDescription(item models.ReceiptItem) int64 {
	points := int64(0)
	factor := constants.ONEFIFTH

	if IsMultipleOf(float64(len(strings.TrimSpace(item.ShortDescription))), float64(constants.THREE)) {

		price, err := strconv.ParseFloat(item.Price, 64)
		if err == nil {
			aux := factor * price
			decimals := aux - math.Round(aux)
			points += int64(math.Round(aux))
			if decimals > 0 && decimals < 0.5 {
				points++
			}
		}

	}

	return points
}

// GetPointsByNumberOfItems calculates points based on the number of receipt items.
func GetPointsByNumberOfItems(totalItems int) int64 {
	return int64((totalItems / 2) * constants.FIVE)
}

// IsMultipleOf checks if a number Is a multiple of another number.
func IsMultipleOf(number float64, multipleOf float64) bool {
	if multipleOf == 0 {
		return false
	}
	quotient := number / multipleOf
	return math.Mod(quotient, 1) == 0
}

// IsDateValue checks if a string represents a valid date.
func IsDateValue(stringDate string) bool {
	_, err := time.Parse("2006-01-02", stringDate)
	return err == nil
}

// IsHourValue checks if a string represents a valid hour.
func IsHourValue(timeString string) bool {

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

// GetPointsByDayFromDateString calculates points based on the day of the purchase date.
func GetPointsByDayFromDateString(date string) int64 {
	points := int64(0)

	dateParsed, err := time.Parse("2006-01-02", date)

	if err == nil && dateParsed.Day()%2 == 1 {
		points = int64(constants.SIX)
	}

	return points
}

// GetPointsByPurchaseTime calculates points based on the purchase time.
func GetPointsByPurchaseTime(purchaseTime string) int64 {
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
