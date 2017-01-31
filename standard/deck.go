package standard

import (
	pc "github.com/xercoy/playcards"
)

// New returns a new, unshuffled deck
func New() *pc.Deck {
	return pc.NewDeck(standardDeck)
}

// NewShuffled returns a new, shuffled deck
func NewShuffled() *pc.Deck {
	return pc.NewShuffledDeck(standardDeck)
}

func generateCards(s pc.Specs) []pc.Card {
	var cards []pc.Card

	for _, suit := range s.Suits {
		for _, rank := range s.Ranks {
			nc, _ := NewCard(suit, rank)
			cards = append(cards, *nc)
		}
	}

	return cards
}
