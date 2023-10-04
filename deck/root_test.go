package deck_test

import (
	"deck"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	deck := deck.New()
	if deck.Len() != 52 {
		t.Errorf("Expected deck length of 52, but got %d", deck.Len())
	}
}

func TestString(t *testing.T) {
	fmt.Println(deck.Card{Suit: deck.Spades, Value: deck.Ace})
	fmt.Println(deck.Card{Suit: deck.Clubs, Value: deck.Nine})
	fmt.Println(deck.Card{Suit: deck.Hearts, Value: deck.Two})
	fmt.Println(deck.Card{Suit: deck.Diamonds, Value: deck.Jack})
	fmt.Println(deck.Card{Suit: deck.Joker})

	// Output:
	// Ace of Spades
	// Nine of Clubs
	// Two of Hearts
	// Jack of Diamonds
	// Joker
}

func TestAddDecks(t *testing.T) {
	deck := deck.New()
	deck.AddDecks(2)
	if deck.Len() != 52*3 {
		t.Errorf("Expected deck length of 156, but got %d", deck.Len())
	}
}

func TestFilter(t *testing.T) {
	d := deck.New()

	cards := []deck.Card{{
		Value: deck.Ace,
	}}
	d.Filter(cards)

	if d.Len() != 52-4 {
		t.Errorf("Expected deck length of 48, but got %d", d.Len())
	}

	d = deck.New()

	cards = []deck.Card{
		{
			Value: deck.Two,
		},
		{
			Value: deck.Two,
		}}
	d.Filter(cards)

	if d.Len() != 52-4 {
		t.Errorf("Expected deck length of 48, but got %d", d.Len())
	}
}
