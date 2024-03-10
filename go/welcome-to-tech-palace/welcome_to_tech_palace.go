package techpalace

import (
	"regexp"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	base := "Welcome to the Tech Palace,"
	return base + " " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" +
		welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	re := regexp.MustCompile(`\*|\n`)
	return strings.Trim(re.ReplaceAllString(oldMsg, ""), " ")
}
