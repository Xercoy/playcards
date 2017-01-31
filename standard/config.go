package standard

import "github.com/xercoy/playcards"

const (
	deckSize = 52
)

var (
	// standard Deck without Jokers
	standardDeck = playcards.Specs{
		Name:        "Standard",
		Suits:       suits,
		Ranks:       ranks,
		Colors:      colors,
		MakeCardsFn: generateCards,
	}
)
