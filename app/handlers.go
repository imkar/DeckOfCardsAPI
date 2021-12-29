package app

import (
	"deckofcards/app/deck"
	"deckofcards/app/models"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Route: "/"
func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Deck of Cards API listening . . .")
	}
}

// TODO:
// CreateDeckHandler -> Saves/Returns Empty DeckOfCards :DONE
// Refactor	:
// Tests

// Route: "/api/createDeck"
func (a *App) CreateDeckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Create New Deck.
		newDeck := new(deck.Deck)
		// Generate new Deck's id.
		deckId := deck.GenerateDeckId()
		deckOfCards := newDeck.CreateNewDeck(deckId)

		// Create Deck model.
		d := &models.Deck{
			ID:               0,
			DeckId:           deckId,
			Shuffled:         false,
			Remaining:        uint8(deckOfCards.GetLength()),
			CreatedDate:      string(time.Now().Format(time.RFC3339)),
			LastModifiedDate: string(time.Now().Format(time.RFC3339)),
		}

		// save to db.
		lastId, err := a.DB.CreateDeck(d)
		if err != nil {
			log.Printf("Could not save into db: v% \n", err)
			//send resp.
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}
		d.ID = lastId
		// map to JSON and send resp.

		dOfc := deckOfCards.GetDeck()
		cards := &models.Cards{
			DeckOfCards: dOfc,
		}

		er := a.DB.CreateCards(cards)
		if er != nil {
			log.Fatal("Cards could not be created")
		}

		resp := mapDeckToJSON(d)
		sendResponse(w, r, resp, http.StatusOK)
	}
}
