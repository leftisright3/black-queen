package decks

import (
	"fmt"
)

type Card struct {
	Suit    Suit
	Rank    Rank
	Visible bool
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", GetRankValueAsString(c.Rank), c.Suit.name)
}

func (c Card) ShortString() string {
	str := fmt.Sprintf("%s-%s", GetRankValueAsString(c.Rank), c.Suit.SingleString())
	return str
}

func (c Card) IsSameCard(to Card) bool {
	if c.Suit.Equals(to.Suit) && c.Rank == to.Rank {
		return true
	}
	return false
}

func (c Card) IsCardSameSuit(to Card) bool {
	if c.Suit.value == to.Suit.value {
		return true
	}
	return false
}

func (c Card) IsSameSuit(to Suit) bool {
	if c.Suit.value == to.value {
		return true
	}
	return false
}

func (c Card) IsSameRank(to Card) bool {
	if c.Rank == to.Rank {
		return true
	}
	return false
}
