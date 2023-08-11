package utils

import (
	"ReceiptProcessor/models"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func CalculateReceiptPoints(receipt *models.Receipt) int64 {
	totalPoints := int64(0)
	totalPoints = getPointsByRetailerName(receipt.Retailer) + getPointsByTotal(receipt.Total) + getPointsByReceiptItems(receipt) + getPointsByPurchaseInfo(receipt.PurchaseDate, receipt.PurchaseTime)
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

func getPointsByReceiptItems(receipt *models.Receipt) int64 {
	points := int64(0)

	println(getPointsByNumberOfItems(len(receipt.Items)))
	points = points + getPointsByNumberOfItems(len(receipt.Items))

	for _, item := range receipt.Items {
		points = points + getPointsByItemDescription(item)
	}

	return points
}

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

func getPointsByItemDescription(item models.ReceiptItem) int64 {
	points := int64(0)
	factor := 0.2
	println(len(strings.TrimSpace(item.ShortDescription)))

	if isMultipleOf(float64(len(strings.TrimSpace(item.ShortDescription))), 3) {

		if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
			auxiliarResult := (factor * price)
			decimals := auxiliarResult - math.Round(auxiliarResult)
			_points := int64(math.Round(auxiliarResult))
			if decimals < 0.5 {
				points = _points + 1
			}
		}

	}

	return points
}

func getPointsByNumberOfItems(totalItems int) int64 {
	return int64((totalItems / 2) * 5)
}

func isMultipleOf(number float64, multipleOf float64) bool {
	if multipleOf == 0 {
		return false
	}
	quotient := number / multipleOf
	return math.Mod(quotient, 1) == 0
}

func isDateValue(stringDate string) bool {
	_, err := time.Parse("2006-01-02", stringDate)
	return err == nil
}

func isHourValue(timeString string) bool {
	isValid := false
	purchaseTime := strings.Split(timeString, ":")
	if hour, err := strconv.Atoi(purchaseTime[0]); err == nil {
		if minute, err := strconv.Atoi(purchaseTime[1]); err == nil {
			println("Hora: " + string(hour))
			println("Minuto: " + string(minute))
			isValid = (hour >= 0 && hour < 24) && (minute >= 0 && minute < 60)
		}
	}
	return isValid
}

func getPointsByDayFromDateString(date string) int64 {
	points := int64(0)
	if dateParsed, err := time.Parse("2006-01-02", date); err == nil {
		fmt.Println(dateParsed.Day())
		/*
		 6 points if the day in the purchase date is odd.
		*/
		if dateParsed.Day()%2 == 1 {
			fmt.Println(dateParsed.Day(), "is Odd number")
			points = int64(6)
		}
	}

	return points
}

func getPointsByPurchaseTime(purchaseTime string) int64 {
	points := int64(0)
	/*
		10 points if the time of purchase is after 2:00pm and before 4:00pm.
	*/
	time := strings.Split(purchaseTime, ":")

	if hour, err := strconv.Atoi(time[0]); err == nil {
		if minute, err := strconv.Atoi(time[1]); err == nil {
			println("Hora: " + string(hour))
			println("Minuto: " + string(minute))
			if (hour >= 14 && hour <= 15) && (minute >= 0 && minute < 60) {
				points = int64(10)
			}
		}
	}

	return points
}
