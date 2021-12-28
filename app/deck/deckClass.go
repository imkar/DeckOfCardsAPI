package deck

import (
	"log"
	"math/rand"
	"time"
)

type Deck struct {
	deckId string
	deck   []Card
}
type Card struct {
	suit  string
	value string
}

var (
	suits  = []string{"HEART", "SPADE", "CLUB", "DIAMOND"}
	values = []string{"ACE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE", "TEN", "JACK", "QUEEN", "KING"}
)

func (d Deck) GetDeck() []Card {
	return d.deck
}

func GenerateDeckId() string {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	fixedLength := 7

	b := make([]rune, fixedLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func generateCard(suit string, val string) Card {
	return Card{
		suit:  suit,
		value: val,
	}
}

func (d Deck) CreateNewDeck(deckId string) Deck {
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			d.deck = append(d.deck, generateCard(suits[i], values[j]))
		}
	}
	d.deckId = deckId
	return d
}

func (d Deck) PrintDeck() {
	for i := 0; i < len(d.deck); i++ {
		log.Printf("SUIT: %v VALUE: %v", d.deck[i].suit, d.deck[i].value)
	}
}

func (d Deck) GetLength() int {
	return len(d.deck)
}
