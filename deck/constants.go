package deck

const (
	Spades Suit = iota
	Hearts
	Clubs
	Diamonds
	Joker
)

const (
	Ace Value = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

var defaultValues = []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

var defaultSuits = []Suit{Spades, Hearts, Clubs, Diamonds}
