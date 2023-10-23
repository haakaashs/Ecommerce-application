package utils

import (
	"regexp"
)

func IsEmailValid(email string) bool {
	// Regular expression for email validation
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
