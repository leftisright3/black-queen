package game

import (
	"blackqueen/decks"
	"sort"
)

type PlayedCardByPlayer struct {
	Player Player
	Card   decks.Card
}

func SortByOrderPosition(pcbp []PlayedCardByPlayer) {
	sort.Slice(pcbp, func(i, j int) bool {
		return pcbp[i].Player.OrderPosition < pcbp[j].Player.OrderPosition
	})
}

func SortBySittingPosition(pcbp []PlayedCardByPlayer) {
	sort.Slice(pcbp, func(i, j int) bool {
		return pcbp[i].Player.SittingPosition < pcbp[j].Player.SittingPosition
	})
}
