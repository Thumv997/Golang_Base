package util

import (
    "strconv"
    "time"
)

// FormatNumber formats a number with a specific precision.
func FormatNumber(number float64, precision int) string {
    return strconv.FormatFloat(number, 'f', precision, 64)
}

// FormatDate formats a time.Time value as a string in the specified format.
func FormatDate(date time.Time, format string) string {
    return date.Format(format)
}

// Other formatting utility functions...
