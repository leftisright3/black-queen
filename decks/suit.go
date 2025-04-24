package decks

type Suit struct {
	value       int
	name        string
	isTrumpSuit bool
}

var (
	Spade = Suit{
		value:       1,
		name:        "Spade",
		isTrumpSuit: false,
	}

	Diamond = Suit{
		value:       2,
		name:        "Diamond",
		isTrumpSuit: false,
	}

	Hearts = Suit{
		value:       3,
		name:        "Hearts",
		isTrumpSuit: false,
	}

	Club = Suit{
		value:       4,
		name:        "Club",
		isTrumpSuit: false,
	}

	suits = [...]Suit{Diamond, Club, Hearts, Spade}
)

func (s Suit) String() string {
	return s.name
}

func (s Suit) Name() string {
	return s.name
}

func (s Suit) Value() int {
	return s.value
}

func (s Suit) Equals(to Suit) bool {
	return s.value == to.value && s.name == to.name
}

func (s Suit) IsTrumpSuit() bool {
	return s.isTrumpSuit
}

func (s Suit) SingleString() string {
	return string(s.name[0])
}
