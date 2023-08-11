package utils

import (
	"ReceiptProcessor/constants"
	"ReceiptProcessor/models"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// CalculateReceiptPoints calculates the total points earned from a receipt.
func CalculateReceiptPoints(receipt *models.Receipt) int64 {
	totalPoints := int64(0)
	totalPoints = getPointsByRetailerName(receipt.Retailer) + getPointsByTotal(receipt.Total) + getPointsByReceiptItems(receipt) + getPointsByPurchaseInfo(receipt.PurchaseDate, receipt.PurchaseTime)
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

	if totalParsed, err := strconv.ParseFloat(total, 64); err == nil {
		fmt.Println(totalParsed)
		if math.Ceil(totalParsed) == totalParsed {
			points = int64(constants.FIFTY)
		}

		if isMultipleOf(totalParsed, constants.ONEQUARTER) {
			points = points + int64(constants.TWENTYFIVE)
		}
	}

	return points
}

// getPointsByReceiptItems calculates points based on receipt items.
func getPointsByReceiptItems(receipt *models.Receipt) int64 {
	points := int64(0)

	println(getPointsByNumberOfItems(len(receipt.Items)))
	points = points + getPointsByNumberOfItems(len(receipt.Items))

	for _, item := range receipt.Items {
		points = points + getPointsByItemDescription(item)
	}

	return points
}

// getPointsByPurchaseInfo calculates points based on purchase date and time.
func getPointsByPurchaseInfo(date string, time string) int64 {
	points := int64(0)
	if isDateValue(date) {
		points = getPointsByDayFromDateString(date)
	}

	if isHourValue(time) {
		points = points + getPointsByPurchaseTime(time)
	}

	return points
}

// getPointsByItemDescription calculates points based on item description.
func getPointsByItemDescription(item models.ReceiptItem) int64 {
	points := int64(0)
	factor := constants.ONEFIFTH
	println(len(strings.TrimSpace(item.ShortDescription)))

	if isMultipleOf(float64(len(strings.TrimSpace(item.ShortDescription))), float64(constants.THREE)) {

		if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
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
	isValid := false
	purchaseTime := strings.Split(timeString, ":")
	if hour, err := strconv.Atoi(purchaseTime[0]); err == nil {
		if minute, err := strconv.Atoi(purchaseTime[1]); err == nil {
			isValid = (hour >= 0 && hour < 24) && (minute >= 0 && minute < 60)
		}
	}
	return isValid
}

// getPointsByDayFromDateString calculates points based on the day of the purchase date.
func getPointsByDayFromDateString(date string) int64 {
	points := int64(0)
	if dateParsed, err := time.Parse("2006-01-02", date); err == nil {
		fmt.Println(dateParsed.Day())
		/*
		 6 points if the day in the purchase date is odd.
		*/
		if dateParsed.Day()%2 == 1 {
			fmt.Println(dateParsed.Day(), "is Odd number")
			points = int64(constants.SIX)
		}
	}

	return points
}

// getPointsByPurchaseTime calculates points based on the purchase time.
func getPointsByPurchaseTime(purchaseTime string) int64 {
	points := int64(0)
	/*
		10 points if the time of purchase is after 2:00pm and before 4:00pm.
	*/
	timeSplit := strings.Split(purchaseTime, ":")

	if hour, err := strconv.Atoi(timeSplit[0]); err == nil {
		if minute, err := strconv.Atoi(timeSplit[1]); err == nil {

			if (hour >= 14 && hour <= 15) && (minute >= 0 && minute < 60) {
				points = int64(constants.TEN)
			}
		}
	}

	return points
}
