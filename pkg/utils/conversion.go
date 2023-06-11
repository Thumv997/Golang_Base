package util

import (
    "strconv"
    "time"
)

// StringToInt converts a string to an integer.
func StringToInt(s string) (int, error) {
    return strconv.Atoi(s)
}

// IntToString converts an integer to a string.
func IntToString(i int) string {
    return strconv.Itoa(i)
}

// FloatToString converts a float64 to a string.
func FloatToString(f float64) string {
    return strconv.FormatFloat(f, 'f', -1, 64)
}

// StringToTime converts a string to a time.Time value using the specified format.
func StringToTime(s string, format string) (time.Time, error) {
    return time.Parse(format, s)
}

// Other conversion functions...