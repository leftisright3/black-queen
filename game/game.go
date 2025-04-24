package game

import (
	"blackqueen/decks"
	"fmt"
	"strings"
)

type Game struct {
	Deck        decks.Deck
	Players     []Player
	Partners    []Player
	NonPartners []Player

	WinningBid        int
	TrumpSuit         decks.Suit
	BaseSuit          decks.Suit
	FirstPartnerCard  decks.Card
	SecondPartnerCard decks.Card

	//PlayedCardsFromPlayers []PlayedCardByPlayer
	WinnerOfHand Player
}

func PrintCardsByPlayer(cardsByPlayer []PlayedCardByPlayer) string {
	var sb strings.Builder
	for _, cardByPlayer := range cardsByPlayer {
		sb.WriteString(fmt.Sprintf("Player: %s = %s ; ", cardByPlayer.Player.Name, cardByPlayer.Card.ShortString()))
	}
	return sb.String()
}

func PrintCardsPlayedByPlayer(players []Player) string {
	var sb strings.Builder
	for _, player := range players {
		sb.WriteString(fmt.Sprintf("Player: %s = %s ; ", player.Name, player.PlayedCard.ShortString()))
	}
	return sb.String()
}
