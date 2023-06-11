package myprivliv
import (
	"math/rand"
	"time"
)

// GenerateRandomString generates a random string of the specified length.
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := len(charset)

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(charsetLength)]
	}

	return string(result)
}

// IsPalindrome checks if the given string is a palindrome.
func IsPalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}