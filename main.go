package main

import (
	"html/template"
	"log"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type Todo struct {
	Id      int
	Message string
}

func main() {

	data := map[string][]Todo{
		"Todos": {
			Todo{Id: 1, Message: "Buy Milk"},
		},
	}

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, data)
	}

	addTodosHandler := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		templ := template.Must(template.ParseFiles("index.html"))
		todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
		//data["Todos"] = append(data["Todos"], todo)
		templ.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo", addTodosHandler)

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
