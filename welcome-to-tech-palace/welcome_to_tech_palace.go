package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// Change customer's name to uppercase.
	customerUpperCase := strings.ToUpper(customer)

	return "Welcome to the Tech Palace, " + customerUpperCase
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// Create stars for up and bottom sides.
	stars := strings.Repeat("*", numStarsPerLine)

	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// Remove all the stars in the message.
	oldMsgWithoutStars := strings.ReplaceAll(oldMsg, "*", "")
	// Remove all the spaces and new lines.
	oldMsgWithoutSpaces := strings.Trim(oldMsgWithoutStars, " \n")

	return oldMsgWithoutSpaces
}
