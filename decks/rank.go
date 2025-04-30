package decks

import "strconv"

type Rank int

const (
	Two     = 2
	Three   = 3
	Four    = 4
	Five    = 5
	Six     = 6
	Seven   = 7
	Eight   = 8
	Nine    = 9
	Ten     = 10
	Jack    = 11
	Queen   = 12
	King    = 13
	Ace     = 14
	minRank = Two
	maxRank = Ace
)

var rankValues = [...]int{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

func GetRankValueAsString(r Rank) string {
	if r >= Two && r <= Ten {
		return strconv.Itoa(int(r))
	} else if r == Jack {
		return "J"
	} else if r == Queen {
		return "Q"
	} else if r == King {
		return "K"
	}
	return "A"
}
