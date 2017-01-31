package playcards

import "math/rand"

// Deck is a data structure that represents a collection of type Card,
// specifications of the details of the card (suits, colors, etc), and
// a function which can generate the slice of cards.
type Deck struct {
	Specs Specs
	cards []Card
}

// NewDeck creates a new Deck in which the elements of the cards field are NOT
// scrambled.
func NewDeck(s Specs) *Deck {
	d := new(Deck)
	d.Specs = s
	d.cards = generateDeckCards(s)

	return d
}

// NewShuffledDeck creates a new Deck in which the elements of the cards field are
// scrambled.
func NewShuffledDeck(s Specs) *Deck {
	d := NewDeck(s)
	d.Shuffle()

	return d
}

// function to generate a standard deck
func generateDeckCards(s Specs) []Card {
	return s.MakeCardsFn(s)
}

// Shuffle randomly rearranges the order of elements of the Card slice of a
// Deck.
func (d *Deck) Shuffle() {
	lock.Lock()
	defer lock.Unlock()

	oldCards := d.cards
	var newCards []Card

	for true {
		// grab a random index
		randNum := rand.Intn(len(oldCards))

		// take the card at that index and add it to the new card collection.
		newCards = append(newCards, oldCards[randNum])

		// remove that card from the copy. If one element is left, break.
		if len(oldCards) > 1 {
			oldCards = append(oldCards[:randNum], oldCards[randNum+1:]...)
		} else {
			break
		}
	}

	d.cards = newCards
}

// Draw returns the card at the end of the Card slice, and finally
// removes the element from the slice. This Method will return true
// if the deck is empty.
func (d *Deck) Draw() (Card, bool) {
	lock.Lock()
	defer lock.Unlock()

	if len(d.cards) == 0 {
		return Card{}, true
	}

	lastElementIndex := len(d.cards) - 1
	drawnCard := d.cards[lastElementIndex]
	d.cards = d.cards[:lastElementIndex]

	return drawnCard, false
}

// Reload discards the current cards with a new, full collection in which
// the elements are NOT randomized.
func (d *Deck) Reload() {
	lock.Lock()
	defer lock.Unlock()

	d.cards = d.Specs.MakeCardsFn(d.Specs)
}

// Cards is a method which returns the collection of cards within an
// instance of a Deck.
func (d *Deck) Cards() []Card {
	return d.cards
}

// ReloadWithShuffle discards the current cards with a new, full collection in which
// the elements are NOT randomized.
func (d *Deck) ReloadWithShuffle() {
	d.Reload()
	d.Shuffle()
}
