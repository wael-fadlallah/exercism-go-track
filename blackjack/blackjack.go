package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var cardsSum = ParseCard(card1) + ParseCard(card2)
	switch {
	case ParseCard(card1) == 11 && ParseCard(card2) == 11:
		return "P"
	case cardsSum == 21 && ParseCard(dealerCard) < 10:
		return "W"
	case cardsSum == 21:
		return "S"
	case cardsSum >= 17 && cardsSum <= 20:
		return "S"
	case cardsSum >= 12 && cardsSum <= 16 && ParseCard(dealerCard) < 7:
		return "S"
	case cardsSum >= 12 && cardsSum <= 16 && ParseCard(dealerCard) < 7:
		return "H"
	case cardsSum <= 11:
		return "H"
	default:
		return "H"
	}
}
