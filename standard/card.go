package standard

import (
	"errors"
	"strings"

	pc "github.com/xercoy/playcards"
)

// standard card logic
var (
	errInvalidCardInput = errors.New("suit or rank input is invalid")

	rankValuePairs = map[string]int{
		"Ace":   1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"10":    10,
		"Jack":  11,
		"Queen": 12,
		"King":  13,
	}

	suitColorPairs = map[string]string{
		"Clubs":    "Black",
		"Spades":   "Black",
		"Diamonds": "Red",
		"Hearts":   "Red",
	}

	ranks = []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	colors = []string{"black", "red"}

	suits = []string{"Clubs", "Spades", "Hearts", "Diamonds"}
)

// NewCard accepts a suit and rank value, validates the input,
// and returns a standard card.
func NewCard(s string, r string) (*pc.Card, error) {
	s = formatString(s)
	r = formatString(r)

	if !ValidateCardFields(s, r) {
		return &pc.Card{}, errInvalidCardInput
	}

	nc := pc.NewCard(s, colorFromSuit(s), valueFromRank(r), r)

	return nc, nil
}

// ValidateCardFields checks the validity of a suit and rank values
// of a card.
func ValidateCardFields(s, r string) bool {
	valid := checkSuit(s)
	if !valid {
		return false
	}

	valid = checkRank(r)
	if !valid {
		return false
	}

	return true
}

func checkRank(r string) bool {
	for _, validRank := range ranks {
		if r == validRank {
			return true
		}
	}
	return false
}

func checkSuit(s string) bool {
	for _, validSuit := range suits {
		if s == validSuit {
			return true
		}
	}
	return false
}

func formatString(str string) string {
	// trim space, convert to lowercase, capitalize first letter
	return strings.Title(strings.ToLower(strings.TrimSpace(str)))
}

func colorFromSuit(s string) string {
	return suitColorPairs[s]
}

func valueFromRank(r string) int {
	return rankValuePairs[r]
}
