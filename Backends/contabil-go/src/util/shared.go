package util

import (
	"math"
	"time"
)

type TagError struct {
	HtmlStatus int
	Inner      error
}

func GetTagError(htmlStatus int, inner error) *TagError {
	return &TagError{HtmlStatus: htmlStatus, Inner: inner}
}

func GetTimeNow() string {
	return time.Now().Format(time.RFC3339)
}

type StringSlice []string

var months = StringSlice{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func (slice StringSlice) IndexOf(element string) int {
	for i, item := range slice {
		if item == element {
			return i
		}
	}
	return -1
}

func MonthToNumber(month string) int {
	return months.IndexOf(month) + 1
}

func MonthSubtractor(month string, year int) (string, int) {

	index := MonthToNumber(month) - 1

	if (index - 1) == 0 {
		m := months[11]
		y := year - 1

		return m, y
	} else {
		m := months[index-1]
		return m, year
	}
}

func MonthSubtractorByJump(month string, year int, jump int) (string, int) {
	for i := 0; i < jump; i++ {
		month, year = MonthSubtractor(month, year)
	}
	return month, year
}

func MonthAdder(month string, year int) (string, int) {
	index := MonthToNumber(month) - 1

	if (index + 1) == 12 {
		month = months[0]
		year = year + 1
	} else {
		month = months[index+1]
	}

	return month, year
}

func NumberToMonth(number int) string {
	return months[number-1]
}

func GetFirstAndLastDayOfCurrentMonth() (string, string) {
	now := time.Now()
	year, month, _ := now.Date()
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1).Add(time.Hour*23 + time.Minute*59 + time.Second*59)
	return firstDay.Format("2006-01-02"), lastDay.Format("2006-01-02T15:04:05")
}

func GetFirstAndLastDayOfMonth(month string, year int) (string, string) {
	monthNumber := MonthToNumber(month)
	firstDay := time.Date(year, time.Month(monthNumber), 1, 0, 0, 0, 0, time.UTC)
	lastDay := firstDay.AddDate(0, 1, -1).Add(time.Hour*23 + time.Minute*59 + time.Second*59)
	return firstDay.Format("2006-01-02"), lastDay.Format("2006-01-02T15:04:05")
}

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	expo := math.Pow(10, float64(precision))
	return float64(Round(num*expo)) / expo
}
