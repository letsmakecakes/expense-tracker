package utils

import "time"

// ParseDate parses a string in the format "YYYY-MM-DD" into time.Time object.
func ParseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

// FormatDate formats a time.Time object into a string in the format "YYYY-MM-DD".
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// IsValidMonth checks if the given month is valid (between 1 and 12).
func IsValidMonth(month int) bool {
	return month >= 1 && month <= 12
}
