package app

import (
	"deckofcards/app/deck"
	"fmt"
	"net/http"
)

//"deckofcards/deck"
//"deckofcards/deck/deckClass"
func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Deck of Cards API listening . . .")
	}
}

func (a *App) CreateDeckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "New deck is being created . . .")
		newDeck := new(deck.Deck)
		deckId := deck.GenerateDeckId()
		deckOfCards := newDeck.CreateNewDeck(deckId)
		deckOfCards.PrintDeck()
		//fmt.Println(deckId)

		// Create Deck
		/*
			d := &models.Deck{
				ID:     0,
				DeckId: ,
				Shuffled: false,
				Remaining:
			}
		*/

	}
}
