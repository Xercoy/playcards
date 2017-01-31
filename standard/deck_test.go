package standard

import (
	"testing"

	"github.com/cheekybits/is"
	pc "github.com/xercoy/playcards"
)

func TestGenerateDeck(t *testing.T) {
	is := is.New(t)
	cardCount := 0
	d := pc.NewShuffledDeck(standardDeck)

	for _ = range d.Cards() {
		cardCount++
	}

	is.Equal(cardCount, deckSize)
}

func TestShuffleDeck(t *testing.T) {
	equal, inequal := 0, 0
	d := pc.NewDeck(standardDeck)
	oldCards := d.Cards()

	d.Shuffle()

	for index, c := range d.Cards() {
		if IsEqual(c, oldCards[index]) {
			equal++
		} else {
			inequal++
		}
	}

	if equal >= inequal {
		t.Errorf("shuffled deck should contain a more random proportion")
	}
}

func IsEqual(c1, c2 pc.Card) bool {
	if (c1.Value() == c2.Value()) && (c1.Suit() == c2.Suit()) {
		return true
	}

	return false
}

func TestDrawCard(t *testing.T) {
	is := is.New(t)
	d := pc.NewShuffledDeck(standardDeck)

	drawnCard, _ := d.Draw()
	t.Logf("Drawn Card: %s", drawnCard.String())

	is.NotEqual(len(d.Cards()), deckSize)
}

func TestDrawUntilEmpty(t *testing.T) {
	is := is.New(t)
	d := pc.NewShuffledDeck(standardDeck)

	drawnCardCount := 0
	for true {
		_, empty := d.Draw()
		if empty {
			break
		} else {
			drawnCardCount++
		}
	}

	is.Equal(drawnCardCount, deckSize)
}

func TestReload(t *testing.T) {
	is := is.New(t)
	d := pc.NewDeck(standardDeck)

	drawnCardCount := 0
	for true {
		_, empty := d.Draw()
		if empty {
			break
		} else {
			drawnCardCount++
		}
	}

	is.Equal(drawnCardCount, deckSize)

	is.Equal(len(d.Cards()), 0)

	d.Reload()

	is.Equal(len(d.Cards()), deckSize)
}

func TestReloadWithShuffle(t *testing.T) {
	is := is.New(t)
	d := pc.NewDeck(standardDeck)

	drawnCardCount := 0
	for true {
		_, empty := d.Draw()
		if empty {
			break
		} else {
			drawnCardCount++
		}
	}

	is.Equal(drawnCardCount, deckSize)

	is.Equal(len(d.Cards()), 0)

	d.ReloadWithShuffle()

	is.Equal(len(d.Cards()), deckSize)
}
