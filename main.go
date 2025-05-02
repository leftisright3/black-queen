package main

import (
	"blackqueen/decks"
	"blackqueen/game"
	"html/template"
	"log"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type NewCard struct {
	CardUrl string
}

func main() {

	var blackqueen game.Game
	blackqueen.Init()
	cardUrl := "https://deckofcardsapi.com/static/img/"

	cardDeckHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html", "partials/card-deck.html", "partials/new-card.html"))
		err := templ.Execute(w, NewCard{CardUrl: "https://deckofcardsapi.com/static/img/AS.png"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	drawCardHandler := func(w http.ResponseWriter, r *http.Request) {
		card := decks.DrawCard(blackqueen.Deck.Cards)
		data := NewCard{
			CardUrl: cardUrl + card.ShortString() + ".png",
		}
		templ := template.Must(template.ParseFiles("partials/new-card.html"))
		err := templ.Execute(w, data)
		//w.Header().Set("Content-Type", "text/html")
		//_, err := w.Write([]byte(``))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	http.HandleFunc("/", cardDeckHandler)
	http.HandleFunc("/new-card", drawCardHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	/**


	var blackqueen game.Game
	blackqueen.Init()
	blackqueen.DealCards()
	blackqueen.TrumpSuit = decks.Spade

	var sb strings.Builder
	for _, player := range blackqueen.Players {
		sb.WriteString(fmt.Sprintf("Player: %s, Sitting Postion: %s, Order Postion: %s | ", player.Name, strconv.Itoa(player.SittingPosition), strconv.Itoa(player.OrderPosition)))
	}
	fmt.Println(sb.String())

	sb.Reset()
	for i, player := range blackqueen.Players {
		player.CardsInHand = decks.SortBySuitAndRankDescending(player.CardsInHand)
		blackqueen.Players[i].CardsInHand = player.CardsInHand
		sb.WriteString(player.Name + "\n" + decks.PrintCards(player.CardsInHand) + fmt.Sprint(len(player.CardsInHand)) + " cards \n")
	}

	fmt.Println(sb.String())

	fmt.Println("\n ***** Playing first card ****** \n")

	game.SortPlayersByOrderPosition(blackqueen.Players)
	for i, player := range blackqueen.Players {
		player.PlayedCard = decks.Card{
			Suit: player.CardsInHand[7].Suit,
			Rank: player.CardsInHand[7].Rank,
		}
		blackqueen.PlayCardFromPlayer(player.PlayedCard, player)
		fmt.Println(player.Name + "\n" + decks.PrintCards(blackqueen.Players[i].CardsInHand) + fmt.Sprint(len(blackqueen.Players[i].CardsInHand)) + " cards\n")
	}

	fmt.Println("Played Cards:")
	fmt.Println(game.PrintCardsPlayedByPlayer(blackqueen.Players) + "\n")

	fmt.Println("Trick Winner:")
	blackqueen.DetermineTrickWinner()
	fmt.Println(blackqueen.WinnerOfHand.Name)

	sb.Reset()
	for _, player := range blackqueen.Players {
		sb.WriteString(fmt.Sprintf("Player: %s, Sitting Postion: %s, Order Postion: %s | ", player.Name, player.SittingPosition))
	}


	*/
}
