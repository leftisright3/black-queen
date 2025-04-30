package decks

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"strings"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	var deck Deck
	deck.Cards = []Card{}
	for _, s := range suits {
		for i := minRank; i <= maxRank; i++ {
			deck.Cards = append(deck.Cards, Card{Suit: s, Rank: Rank(i)})
		}
	}
	return deck
}

func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(cards))
	shuffledCards := make([]Card, len(cards))
	for i, randIndex := range perm {
		shuffledCards[i] = cards[randIndex]
	}
	copy(cards, shuffledCards)
	return cards
}

func DrawCard(cards []Card) Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	idx := r.Intn(52)
	return cards[idx]
}

func SortBySuitAndRankAscending(cards []Card) []Card {
	diamonds := FilterOutBySuit(cards, Diamond)
	diamonds = SortByRankAscending(diamonds)

	clubs := FilterOutBySuit(cards, Club)
	clubs = SortByRankAscending(clubs)

	hearts := FilterOutBySuit(cards, Hearts)
	hearts = SortByRankAscending(hearts)

	spades := FilterOutBySuit(cards, Spade)
	spades = SortByRankAscending(spades)

	cards = slices.Concat(diamonds, clubs, hearts, spades)
	return cards
}

func SortBySuitAndRankDescending(cards []Card) []Card {
	diamonds := FilterOutBySuit(cards, Diamond)
	diamonds = SortByRankDescending(diamonds)

	clubs := FilterOutBySuit(cards, Club)
	clubs = SortByRankDescending(clubs)

	hearts := FilterOutBySuit(cards, Hearts)
	hearts = SortByRankDescending(hearts)

	spades := FilterOutBySuit(cards, Spade)
	spades = SortByRankDescending(spades)

	cards = slices.Concat(diamonds, clubs, hearts, spades)
	return cards
}

func SortByRankAscending(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank < cards[j].Rank
	})
	return cards
}

func SortByRankDescending(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Rank > cards[j].Rank
	})
	return cards
}

func FilterOutBySuit(cards []Card, suit Suit) []Card {
	var filteredCards []Card
	for _, card := range cards {
		if card.Suit.Equals(suit) {
			filteredCards = append(filteredCards, card)
		}
	}
	return filteredCards
}

func FilterOutByRank(cards []Card, r Rank) []Card {
	for i, c := range cards {
		if c.Rank == r {
			cards = append(cards[:i], cards[i+1:]...)
		}
	}
	return cards
}

func (deck *Deck) Print() string {
	var sb strings.Builder
	for _, card := range deck.Cards {
		sb.WriteString(card.ShortString() + "\n")
	}
	return sb.String()
}

func PrintCards(cards []Card) string {
	var sb strings.Builder
	for _, card := range cards {
		sb.WriteString(fmt.Sprintf("%s, ", card.ShortString()))
	}
	return sb.String()
}
