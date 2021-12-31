package app

import (
	"deckofcards/app/deck"
	"deckofcards/app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Parses incoming request to provided data model interface.
func requestParser(w http.ResponseWriter, r *http.Request, data interface{}) error {
	fmt.Println(json.NewDecoder(r.Body).Decode(data))
	return json.NewDecoder(r.Body).Decode(data)
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Printf("Cannot format JSON. err=%v\n", err)
	}
}

// Maps Deck db model to Json.
func mapDeckToJSON(d *models.Deck) models.JsonDeck {
	return models.JsonDeck{
		ID:               d.ID,
		DeckId:           d.DeckId,
		Shuffled:         d.Shuffled,
		Remaining:        d.Remaining,
		CreatedDate:      d.CreatedDate,
		LastModifiedDate: d.LastModifiedDate,
	}
}

func mapCardToJSON(c *deck.Card) models.JsonCard {
	return models.JsonCard{
		Suit:  c.Suit,
		Value: c.Value,
	}
}

func mapDrawCardToJSON(jd *models.JsonDeck, jc *models.JsonCard) models.JsonDraw {
	return models.JsonDraw{
		DeckJson: *jd,
		CardJson: *jc,
	}
}
