package dates

import (
	"math"
	"time"
)

// DateFormat represents the parsing format for a date string
// TimeFormat represents the parsing format for a time string
// DateTimeFormat represents the parsing format for a date time string
const (
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
	DateTimeFormat = "2006-01-02 15:04:05"
)

// DateStringToTime parses a date string to a Time struct
func DateStringToTime(date string) (time.Time, error) {
	return time.Parse(DateFormat, date)
}

// DateTimeStringToTime parses a date time string to a Time struct
func DateTimeStringToTime(date string) (time.Time, error) {
	return time.Parse(DateTimeFormat, date)
}

// GetDateThisMorning sets the time for today at midnight (in the morning, not evening)
func GetDateThisMorning() time.Time {
	return SetTimeToBeginDay(time.Now())
}

// SetTimeToNoon sets the given date to 00:00:00.0
func SetTimeToNoon(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 12, 0, 0, 0, time.UTC)
}

// SetTimeToBeginDay sets the given date to 00:00:00.0
func SetTimeToBeginDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}

// SetTimeToEndDay sets the given date to 23:59:59.99
func SetTimeToEndDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 99, time.UTC)
}

// TimeStringFromTime formats a Time object to a time string
func TimeStringFromTime(date time.Time) string {
	return date.Format(TimeFormat)
}

// DateStringFromTime formats a Time object to a time string
func DateStringFromTime(date time.Time) string {
	return date.Format(DateFormat)
}

// DateTimeStringFromTime formats a Time object to a date time string
func DateTimeStringFromTime(date time.Time) string {
	return date.Format(DateTimeFormat)
}

// DifferenceInDays gets the difference of the given dates in days
func DifferenceInDays(startDate time.Time, endDate time.Time) int {
	duration := endDate.Sub(startDate)

	return roundFloatToInt(duration.Seconds() / 86400)
}

// DifferenceInHours gets the difference of the given dates in hours
func DifferenceInHours(startDate time.Time, endDate time.Time) int {
	duration := endDate.Sub(startDate)

	return roundFloatToInt(duration.Hours())
}

// roundFloatToInt rounds off any floats to integers
func roundFloatToInt(input float64) int {
	var result float64
	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}
	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)
	return int(i)
}

// IsSameOrBefore determines whether the first date is equal or before the second date
func IsSameOrBefore(date time.Time, comparison time.Time) bool {
	return date.Equal(comparison) || date.Before(comparison)
}

// IsSameOrAfter determines whether the first date is equal or after the second date
func IsSameOrAfter(date time.Time, comparison time.Time) bool {
	return date.Equal(comparison) || date.After(comparison)
}


func DayStartUnix(year, month, day int) int64 {
	// Defining t in UTC for Unix method
	t := time.Date(year, time.Month(month), day, 0, 0, 1, 0, time.Now().Location())
	// Calling Unix method
	unix := t.Unix()
	return unix
}


func DayEndUnix(year, month, day int) int64 {
	// Defining t in UTC for Unix method
	t := time.Date(year, time.Month(month), day, 23, 59, 59, 0, time.Now().Location())
	// Calling Unix method
	unix := t.Unix()
	return unix
}

func YesterdayDates() (int, int, int) {
	yesterday := time.Now().AddDate(0, 0, -1)
	return yesterday.Year(), int(yesterday.Month()), yesterday.Day()
}