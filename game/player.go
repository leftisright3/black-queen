package game

import (
	"blackqueen/decks"
	"sort"
)

type Player struct {
	Name            string
	CardsInHand     []decks.Card
	WonHands        []decks.Card
	SittingPosition int
	OrderPosition   int
	PlayedCard      decks.Card
	HasPlayedCard   bool
}

func SortPlayersByOrderPosition(player []Player) {
	sort.Slice(player, func(i, j int) bool {
		return player[i].OrderPosition < player[j].OrderPosition
	})
}

func SortPlayersBySittingPosition(player []Player) {
	sort.Slice(player, func(i, j int) bool {
		return player[i].SittingPosition < player[j].SittingPosition
	})
}
