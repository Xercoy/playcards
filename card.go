package playcards

import "fmt"

// Card represents a data structure for a standard playing card.
type Card struct {
	suit  string // the suit of the card
	color string // the color of the card
	value int    // the card's number value
	rank  string // the card's face value
}

// String is a simple string representation of the card's rank and suit.
func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank(), c.Suit())
}

// Suit is a string value which returns one of the four possible suits.
func (c *Card) Suit() string {
	return c.suit
}

// Color returns a card's color
func (c *Card) Color() string {
	return c.color
}

// Value returns a card's number value.
func (c *Card) Value() int {
	return c.value
}

// Rank returns a card's rank.
func (c *Card) Rank() string {
	return c.rank
}

// NewCard takes a suit, color, and rank as strings and
// returns a copy of a new Card.
func NewCard(s string, c string, v int, r string) *Card {
	return &Card{
		suit:  s,
		color: c,
		value: v,
		rank:  r,
	}
}
