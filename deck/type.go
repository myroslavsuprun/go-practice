package deck

type Suit uint8

type Value uint8

type Card struct {
	Suit  Suit
	Value Value
}

type Deck []Card

type SortFunc func(i, j int) bool

type IDeck interface {
	Sort(f SortFunc)
	Shuffle()
	AddJokers(n int)
	Filter(f []Card)
	AddDecks(n int)
	Less(i, j int) bool
	Len() int
}
