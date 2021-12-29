package deck

import (
	"log"
	"math/rand"
	"time"
)

type Deck struct {
	DeckId      string
	DeckOfCards Cards
}

type Cards []Card
type Card struct {
	Suit  string
	Value string
}

var (
	suits  = []string{"HEART", "SPADE", "CLUB", "DIAMOND"}
	values = []string{"ACE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE", "TEN", "JACK", "QUEEN", "KING"}
)

func (d Deck) GetDeck() []Card {
	return d.DeckOfCards
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
		Suit:  suit,
		Value: val,
	}
}

func (d Deck) CreateNewDeck(deckId string) Deck {
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			d.DeckOfCards = append(d.DeckOfCards, generateCard(suits[i], values[j]))
		}
	}
	d.DeckId = deckId
	return d
}

func (d Deck) PrintDeck() {
	for i := 0; i < len(d.DeckOfCards); i++ {
		log.Printf("SUIT: %v VALUE: %v", d.DeckOfCards[i].Suit, d.DeckOfCards[i].Value)
	}
}

func (d Deck) GetLength() int {
	return len(d.DeckOfCards)
}
