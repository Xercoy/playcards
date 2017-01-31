package standard

import (
	"testing"

	"github.com/cheekybits/is"
)

func TestMethods(t *testing.T) {
	is := is.New(t)

	c, err := NewCard("SpaDES", "10")
	is.NoErr(err)

	is.Equal(c.Color(), "Black")
	is.Equal(c.Rank(), "10")
	is.Equal(c.Value(), 10)
	is.Equal(c.Suit(), "Spades")
	is.Equal(c.String(), "10 of Spades")
}

func TestValidation(t *testing.T) {
	is := is.New(t)

	c1, err := NewCard("Diamonds", "5")
	is.NoErr(err)

	valid := ValidateCardFields(c1.Suit(), c1.Rank())
	is.True(valid)

	c2, err := NewCard("Monkeys", "Jack")
	is.Err(err)

	valid = ValidateCardFields(c2.Suit(), c2.Rank())
}

func TestValidCardSuit(t *testing.T) {
	is := is.New(t)

	_, err := NewCard("spades", "10")
	is.NoErr(err)
}

func TestInvalidCardSuit(t *testing.T) {
	is := is.New(t)

	_, err := NewCard("spears", "Jack")
	is.Err(err)
}

func TestValidCardRank(t *testing.T) {
	is := is.New(t)

	_, err := NewCard("Hearts", "5")
	is.NoErr(err)
}

func TestInvalidCardRank(t *testing.T) {
	is := is.New(t)

	_, err := NewCard("Clubs", "Banana")
	is.Err(err)
}
