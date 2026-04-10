package techpalace

import (
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    stars := strings.Repeat("*",numStarsPerLine)
    addBorder := stars+"\n"+welcomeMsg+"\n"+stars
    return addBorder
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	removeStars := strings.ReplaceAll(oldMsg, "*", "")
    removeSpace := strings.TrimSpace(removeStars)
    return removeSpace
	panic("Please implement the CleanupMessage() function")
}
