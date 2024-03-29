// Code generated by "stringer -type=Suit,Value"; DO NOT EDIT.

package deck

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Spades-0]
	_ = x[Hearts-1]
	_ = x[Clubs-2]
	_ = x[Diamonds-3]
	_ = x[Joker-4]
}

const _Suit_name = "SpadesHeartsClubsDiamondsJoker"

var _Suit_index = [...]uint8{0, 6, 12, 17, 25, 30}

func (i Suit) String() string {
	if i >= Suit(len(_Suit_index)-1) {
		return "Suit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Suit_name[_Suit_index[i]:_Suit_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ace-0]
	_ = x[Two-1]
	_ = x[Three-2]
	_ = x[Four-3]
	_ = x[Five-4]
	_ = x[Six-5]
	_ = x[Seven-6]
	_ = x[Eight-7]
	_ = x[Nine-8]
	_ = x[Ten-9]
	_ = x[Jack-10]
	_ = x[Queen-11]
	_ = x[King-12]
}

const _Value_name = "AceTwoThreeFourFiveSixSevenEightNineTenJackQueenKing"

var _Value_index = [...]uint8{0, 3, 6, 11, 15, 19, 22, 27, 32, 36, 39, 43, 48, 52}

func (i Value) String() string {
	if i >= Value(len(_Value_index)-1) {
		return "Value(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Value_name[_Value_index[i]:_Value_index[i+1]]
}
