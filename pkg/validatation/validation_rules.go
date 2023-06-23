package validation

import "regexp"

// ValidateUsername checks if a username is valid based on certain criteria.
func ValidateUsername(username string) bool {
    // Perform validation logic here
    // Return true if the username is valid, false otherwise
	return true
}

// ValidateEmail checks if an email address is valid based on certain criteria.
func ValidateEmail(email string) bool {
	// Use regular expression to validate email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
    // Perform validation logic here
    // Return true if the email is valid, false otherwise
}

// ValidatePassword checks the password against predefined rules
func ValidatePassword(password string) bool {
	// Define rules for password validation
	// Example rules: minimum length of 8 characters, at least one uppercase letter, one lowercase letter, and one digit

	if len(password) < 8 {
		return false
	}

	hasUppercase := false
	hasLowercase := false
	hasDigit := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUppercase = true
		} else if char >= 'a' && char <= 'z' {
			hasLowercase = true
		} else if char >= '0' && char <= '9' {
			hasDigit = true
		}
	}

	return hasUppercase && hasLowercase && hasDigit
}


// Other input validation functions...
