package game

import (
	"blackqueen/decks"
	"slices"
)

func (g *Game) Init() {
	g.Deck = decks.NewDeck()
	g.Players = append(g.Players, CreateNewPlayer("Kushan", 0, 2))
	g.Players = append(g.Players, CreateNewPlayer("Rohan", 1, 3))
	g.Players = append(g.Players, CreateNewPlayer("Tiree", 2, 0))
	g.Players = append(g.Players, CreateNewPlayer("Madison", 3, 1))
}

func CreateNewPlayer(name string, sittingPosition int, orderPosition int) Player {
	player := Player{
		Name:            name,
		SittingPosition: sittingPosition,
		OrderPosition:   orderPosition,
		HasPlayedCard:   false,
	}
	return player
}

func (g *Game) DealCards() {
	g.Deck.Cards = decks.Shuffle(g.Deck.Cards)
	cardsPerPlayer := len(g.Deck.Cards) / len(g.Players)
	dealEndIndex := cardsPerPlayer
	dealStartIndex := 0
	for i, player := range g.Players {
		for dealStartIndex <= (dealEndIndex - 1) {
			player.CardsInHand = append(player.CardsInHand, g.Deck.Cards[dealStartIndex])
			dealStartIndex++
		}
		dealEndIndex += cardsPerPlayer
		g.Players[i] = player
	}
}

func (g *Game) PlayCardFromPlayer(playedCard decks.Card, player Player) {
	// need to set player's PlayedCard and HasPlayed properties in this function
	player.PlayedCard = playedCard
	for i, card := range player.CardsInHand {
		if playedCard.IsSameCard(card) && player.HasPlayedCard == false {
			player.CardsInHand = slices.Delete(player.CardsInHand, i, i+1)
			player.HasPlayedCard = true
		}
	}

	for i, p := range g.Players {
		if p.Name == player.Name {
			g.Players[i] = player
		}
	}
}

/**
func (g *Game) DetermineTrickWinner() {
	SortByOrderPosition(g.PlayedCardsFromPlayers)
	var highestBaseRank decks.Rank
	var highestTrumpRank decks.Rank
	if g.PlayedCardsFromPlayers[0].Player.OrderPosition == 0 {
		g.BaseSuit = g.PlayedCardsFromPlayers[0].Card.Suit
		highestBaseRank = g.PlayedCardsFromPlayers[0].Card.Rank
	}

	var winningPlayerIdx int
	for i, pcfp := range g.PlayedCardsFromPlayers[1:] {
		if pcfp.Card.IsSameSuit(g.BaseSuit) {
			if pcfp.Card.Rank > highestBaseRank {
				highestBaseRank = pcfp.Card.Rank
				g.WinnerOfHand = pcfp.Player
			}
		} else if pcfp.Card.IsSameSuit(g.TrumpSuit) {
			if pcfp.Card.Rank > highestTrumpRank {
				highestTrumpRank = pcfp.Card.Rank
				g.WinnerOfHand = pcfp.Player
				winningPlayerIdx = i
			}
		} else {
			g.WinnerOfHand = pcfp.Player
		}
	}
	g.SetNewOrderPosition(winningPlayerIdx)
}
 **/

func (g *Game) DetermineTrickWinner() {
	SortPlayersByOrderPosition(g.Players)
	var highestBaseRank decks.Rank
	var highestTrumpRank decks.Rank

	// first set base rank from the player in the 0th order position
	if g.Players[0].OrderPosition == 0 {
		g.BaseSuit = g.Players[0].PlayedCard.Suit
		highestBaseRank = g.Players[0].PlayedCard.Rank
		// set them as the winner for now..
		g.WinnerOfHand = g.Players[0]
	}

	var winningPlayerIdx int
	// iterate over the rest of the players from the 1st order position
	for i, player := range g.Players[1:] {
		if player.PlayedCard.IsSameSuit(g.BaseSuit) {
			if player.PlayedCard.Rank > highestBaseRank {
				highestBaseRank = player.PlayedCard.Rank
				// update the winner
				g.WinnerOfHand = player
				winningPlayerIdx = i
			}
		} else {
			if player.PlayedCard.IsSameSuit(g.TrumpSuit) {
				if player.PlayedCard.Rank > highestTrumpRank {
					highestTrumpRank = player.PlayedCard.Rank
					// update the winner
					g.WinnerOfHand = player
					winningPlayerIdx = i
				}
			}
		}
	}

	// remember to re-sort the players array by the winning player
	g.SetNewOrderPosition(winningPlayerIdx)
}

func (g *Game) SetNewOrderPosition(winningPlayerIdx int) {
	newOrderPostion := 0
	for i := winningPlayerIdx; i < len(g.Players); i++ {
		g.Players[i].OrderPosition = newOrderPostion
		newOrderPostion++
	}
	for i := 0; i <= winningPlayerIdx; i++ {
		g.Players[i].OrderPosition = newOrderPostion
	}
}
