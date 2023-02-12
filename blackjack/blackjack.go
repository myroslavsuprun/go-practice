package blackjack

// Predeclared strategies of the game for the player.
const stand = "S"
const hit = "H"
const split = "P"
const win = "W"

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// I know this code is very poor quality, sorry.
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerNumber := ParseCard(card1) + ParseCard(card2)
	dealerNumber := ParseCard(dealerCard)

	switch {
	case playerNumber == 22:
		return split
	case playerNumber == 21:
		if dealerNumber >= 10 {
			return stand
		}
		return win
	case playerNumber >= 17 && playerNumber <= 20:
		return stand
	case playerNumber >= 12 && playerNumber <= 16:
		if dealerNumber >= 7 {
			return hit
		}
		return stand
	default:
		return hit
	}
}
