package app

import (
	"deckofcards/app/deck"
	"deckofcards/app/models"
	"fmt"
	"log"
	"math/rand"
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

		if err := a.DB.CreateCards(cards); err != nil {
			log.Fatalf("Cards could not be created. Err: %v", err)
		}

		resp := mapDeckToJSON(d)
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func draw(cardsObj deck.Cards) (deck.Card, deck.Cards) {
	return cardsObj[0], cardsObj[1:]
}

// TODO:
// - NEEDS REFACTOR -> This function must be well designed.
//	|-> Check for aggregator pattern.
//  |-> Encapsulation
// - NEEDS TESTING
// - GO ROUTINE MAYBE ?

func (a *App) DrawCardByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get Params.
		vars := mux.Vars(r)
		deckid := vars["deckid"]

		// CALL getCardsByIdHandler
		cardsObj := a.getCardsByIdHandler(deckid)

		// GET FIRST CARD! AND REMOVE FIRST CARD FROM DB.
		var drawnCard deck.Card
		var restOfCards deck.Cards

		if len(cardsObj) > 1 {

			// Draw a card.
			drawnCard, restOfCards = draw(cardsObj)

			// CALL decrementRemainingByDeckIdHandler()
			if err := a.decrementRemainingByDeckIdHandler(deckid); err != nil {
				log.Fatalf("Cannot decrement the count of cards. Err: %v", err)
			}

			// CALL UPDATE the REST
			if err := a.updateCardsByDeckId(restOfCards, deckid); err != nil {
				log.Fatalf("Could not update the rest of the cards by id. Err: %v", err)
			}

			d := a.getDeckByIdHandler(deckid)
			dm := &models.Deck{
				ID:               d.ID,
				DeckId:           d.DeckId,
				Shuffled:         d.Shuffled,
				Remaining:        d.Remaining,
				CreatedDate:      d.CreatedDate,
				LastModifiedDate: d.LastModifiedDate,
			}
			jsonDeck := mapDeckToJSON(dm)
			jsonCard := mapCardToJSON(&drawnCard)
			jsonDraw := mapDrawCardToJSON(&jsonDeck, &jsonCard)
			sendResponse(w, r, jsonDraw, http.StatusOK)

		} else if len(cardsObj) == 1 {

			drawnCard := cardsObj[0]

			// CALL decrementRemainingByDeckIdHandler()
			if err := a.decrementRemainingByDeckIdHandler(deckid); err != nil {
				log.Fatalf("Cannot decrement the count of cards. Err: %v", err)
			}

			// CALL UPDATE the REST
			if err := a.updateCardsByDeckId(restOfCards, deckid); err != nil {
				log.Fatalf("Could not update the rest of the cards by id. Err: %v", err)
			}

			d := a.getDeckByIdHandler(deckid)
			dm := &models.Deck{
				ID:               d.ID,
				DeckId:           d.DeckId,
				Shuffled:         d.Shuffled,
				Remaining:        d.Remaining,
				CreatedDate:      d.CreatedDate,
				LastModifiedDate: d.LastModifiedDate,
			}
			jsonDeck := mapDeckToJSON(dm)
			jsonCard := mapCardToJSON(&drawnCard)
			jsonDraw := mapDrawCardToJSON(&jsonDeck, &jsonCard)
			sendResponse(w, r, jsonDraw, http.StatusOK)

			// CLEAN deckId from DB cards and decks.
			//DeleteCardsByDeckId(deckid)
			if err := a.DeleteCardsByDeckId(deckid); err != nil {
				log.Fatalf("Cards cannot be deleted from db: %v", err)
			}
			//DeleteDeckByDeckId(deckid)
			if err := a.DeleteDeckByDeckId(deckid); err != nil {
				log.Fatalf("Decks cannot be deleted from db: %v", err)
			}

		} else {
			log.Println("No cards left in the deck.")
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}
	}
}

func (a *App) DeleteCardsByDeckId(deckid string) error {
	err := a.DB.DeleteCardsByDeckId(deckid)
	return err
}

func (a *App) DeleteDeckByDeckId(deckid string) error {
	err := a.DB.DeleteDeckByDeckId(deckid)
	return err
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
		log.Fatalf("Could not decrement the Remaining, err: %v", err)
	}
	return err
}

func (a *App) updateCardsByDeckId(dOfC deck.Cards, deckid string) error {
	err := a.DB.UpdateCardsByDeckId(dOfC, deckid)
	if err != nil {
		log.Fatalf("Could not update the rest of the deck after card drawn. %v", err)
	}
	return err
}

func (a *App) getDeckByIdHandler(deckid string) models.Deck {
	decksObj := a.DB.GetDeckByDeckId(deckid)
	return decksObj
}

func shuffle(cardsObj deck.Cards) deck.Cards {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cardsObj), func(i, j int) { cardsObj[i], cardsObj[j] = cardsObj[j], cardsObj[i] })
	return cardsObj
}

func (a *App) ShuffleDeckByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		deckid := params["deckid"]
		// Get cards from DB by deckId.
		cardsObj := a.getCardsByIdHandler(deckid)
		// Shuffle cards
		cardsObj = shuffle(cardsObj)
		// Update DB with shuffled cards.
		err := a.updateCardsByDeckId(cardsObj, deckid)
		if err != nil {
			log.Fatalf("Could not update the rest of the deck after card drawn. %v", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}
		sendResponse(w, r, &models.Status{DeckId: deckid, Status: 200}, http.StatusOK)
	}
}
