package techpalace

import (
	"fmt"
	"strings"
)


// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	capital := strings.ToUpper(customer)
	return fmt.Sprintf("Welcome to the Tech Palace, %s", capital)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := ""
	for i := 0; i < numStarsPerLine; i++ {
		border += "*"
	}
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	lines := strings.Split(oldMsg, "\n")
	if len(lines) < 2 {
		return ""
	}
	line := lines[1]
	clean := strings.Trim(line, "*")
	clean = strings.TrimSpace(clean)
	return clean
}
