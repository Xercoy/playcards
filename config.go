package playcards

import (
	"math/rand"
	"sync"
	"time"
)

var (
	// protects shuffle, draw and reload
	lock sync.RWMutex 
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano()) // set seed for shuffling
}

// Specs are the details used to create a deck of a particular type.
type Specs struct {
	Name        string
	Suits       []string
	Ranks       []string
	Colors      []string
	MakeCardsFn func(Specs) []Card
}
