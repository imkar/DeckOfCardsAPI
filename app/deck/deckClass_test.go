package deck

import (
	"fmt"
	"regexp"
	"testing"
)

func TestCreateNewDeck(t *testing.T) {
	// Initialize deck obj.
	deck := new(Deck)
	// create Deck id.
	deckid := GenerateDeckId()
	// create new Deck
	newDeck := deck.CreateNewDeck(deckid)

	// Is deck id is equal to newly created decks id ?
	if newDeck.DeckId != deckid {
		t.Fatalf("Created deck id must be equal to %v, instead %v", deckid, newDeck.DeckId)
	}

}

func TestGetLength(t *testing.T) {
	// Initialize deck obj.
	deck := new(Deck)
	// create Deck id.
	deckid := GenerateDeckId()
	// create new Deck
	newDeck := deck.CreateNewDeck(deckid)
	// Does newly created deck has 52 cards ?
	if newDeck.GetLength() != 52 {
		t.Fatalf("Created deck must have 52 CARDS, not %v", len(newDeck.DeckOfCards))
	}
}

func TestGenerateDeckId(t *testing.T) {
	// create Deck id.
	deckid := GenerateDeckId()
	// REGEXP for deckId.
	r, _ := regexp.Compile("[a-zA-Z]{7}")
	// Does compiled regex comprises deckId generated?
	if !r.MatchString(deckid) {
		t.Fatalf("Deck id does not match the limitations: %v", deckid)
	}
}

// returns type of a variable as string.
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
func TestGetDeck(t *testing.T) {
	// Initialize deck obj.
	deck := new(Deck)
	// create Deck id.
	deckid := GenerateDeckId()
	// create new Deck
	newDeck := deck.CreateNewDeck(deckid)
	typeVar := typeof(newDeck.GetDeck())
	if "deck.Cards" != typeVar {
		t.Fatalf("GetDeck() does not return type of 'Cards', instead returns %v", typeVar)
	}
}

func TestPrintDeck(t *testing.T) {
	// Initialize deck obj.
	deck := new(Deck)
	// create Deck id.
	deckid := GenerateDeckId()
	// create new Deck
	newDeck := deck.CreateNewDeck(deckid)
	// Prints deck
	newDeck.PrintDeck()
}
