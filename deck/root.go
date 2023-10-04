// Create a deck of cards
package deck

import (
	"fmt"
	"math/rand"
)

/*
Create a new deck of cards;

Initial legth of the deck is 52 with the following order:

Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King;
Spades, Hearts, Clubs, Diamonds.
*/
func New() IDeck {
	var cards Deck
	for _, s := range defaultSuits {
		for _, v := range defaultValues {
			cards = append(cards, Card{Suit: s, Value: v})
		}
	}
	return &cards
}

func (d *Deck) Len() int {
	return len(*d)
}

// Sort sorts the deck using the provided SortFunc.
func (d *Deck) Sort(f SortFunc) {
	for i := range *d {
		for j := range *d {
			if f(i, j) {
				(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
			}
		}
	}
}

// Shuffle shuffles the deck.
func (d *Deck) Shuffle() {
	for i := range *d {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}

// AddJokers adds n amount of jokers to the deck.
func (d *Deck) AddJokers(n int) {
	for i := 0; i < n; i++ {
		*d = append(*d, Card{Suit: Joker})
	}
}

// Filter removes the provided cards from the deck.
func (d *Deck) Filter(f []Card) {
	for _, fCard := range f {
		for i, dCard := range *d {
			if fCard.Value == dCard.Value {
				*d = append((*d)[:i], (*d)[i+1:]...)
			}
		}
	}
}

// Less is a default SortFunc that sorts the deck by card value.
func (d *Deck) Less(i, j int) bool {
	return (*d)[i].Value < (*d)[j].Value
}

// AddDecks adds n amount of decks to the deck;
// added decks are copies of the current one.
func (d *Deck) AddDecks(n int) {
	c := *d
	for i := 0; i < n; i++ {
		*d = append(*d, c...)
	}
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	return fmt.Sprintf("%s of %ss", c.Value.String(), c.Suit.String())
}
