package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "one":
		return 1
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
	case "ten",
		"jack",
		"queen",
		"king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	hand := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case hand > 21:
		return "P"
	case hand == 21 && dealer < 10:
		return "W"
	case hand == 21 && dealer >= 10 ||
		hand >= 17 && hand <= 20 ||
		hand >= 12 && hand <= 16 && dealer < 7:
		return "S"
	default:
		return "H"
	}
}
