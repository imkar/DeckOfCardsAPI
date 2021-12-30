package app

import (
	"deckofcards/app/deck"
	"deckofcards/app/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
		// Get Params.
		vars := mux.Vars(r)

		//fmt.Printf("Deck returned: %v\n", decksObj)
		// CALL getCardsByIdHandler
		cardsObj := a.getCardsByIdHandler(vars["deckid"])

		// GET FIRST CARD! AND REMOVE FIRST CARD FROM DB.
		var drawnCard deck.Card
		if len(cardsObj) > 0 {
			drawnCard = cardsObj[0]
			// CALL decrementRemainingByDeckIdHandler()
			err := a.decrementRemainingByDeckIdHandler(vars["deckid"])
			if err != nil {
				log.Fatalf("Cannot decrement the count of cards.")
			}

			// CALL UPDATE WITH REST
		} else {
			log.Println("No cards left in the deck.")
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}
		fmt.Printf("Drawn Cards is %v of %v", drawnCard.Value, drawnCard.Suit)
		// UPDATE DECK STATE ON DB.

		// RETURN THE DRAWN CARD.
	}
}

// This should be helper or someting different...
func (a *App) getCardsByIdHandler(deckid string) deck.Cards {
	// TODO: (this is an internal function),
	//		(can be written as interface),
	//		(check carefully the input, this function is an internal executor.)
	cardsObj := a.DB.GetCardsByDeckId(deckid)
	return cardsObj
}

func (a *App) decrementRemainingByDeckIdHandler(deckid string) error {
	err := a.DB.DecrementRemainingCountOnDeckById(deckid)
	if err != nil {
		log.Fatalf("Could not decrement the Remaining")
	}
	return err
}

/*
func (a *App) getDeckByIdHandler(deckid string) {
	a.DB.GetDeckByDeckId(deckid)
	//return decksObj
}
*/
