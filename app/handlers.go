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
// Refactor	:DONE
// Tests :DONE
// if deckId exists, checks it (DeckID must be UNIQUE).

// Route: "/api/createDeck"
func (a *App) CreateDeckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Initializes Deck type.
		newDeck := new(deck.Deck)
		// Generate new Deck's id.
		deckId := deck.GenerateDeckId()
		// Creates deck of cards.
		deckOfCards := newDeck.CreateNewDeck(deckId)

		// Fill Deck model for DB.
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

		cards := &models.Doc{
			DeckId:      deckId,
			DeckOfCards: deckOfCards.GetDeck(),
		}

		er := a.DB.CreateCards(cards)
		if er != nil {
			log.Fatal("Cards could not be created")
		}

		resp := mapDeckToJSON(d)
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) DrawCardByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO:

		// CALL getDeckByIdHandler()

		// GET FIRST CARD! AND REMOVE FIRST CARD FROM DB.

		// UPDATE DECK STATE ON DB.

		// RETURN THE DRAWN CARD.
	}
}

// This should be helper or someting different...

func getDeckByIdHandler() {
	// TODO: (this is an internal function),
	//		(can be written as interface),
	//		(check carefully the input, this function is an internal executor.)

	// get deckid from params.

	// check whether these deckid exists in the db.

	// get deck from db by id.
	//// a.DB.GetByDeckId(params["deckid"])

	// unmarshall json to struct.

	// print deck

}
