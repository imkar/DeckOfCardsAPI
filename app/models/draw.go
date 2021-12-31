package models

type JsonCard struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

type JsonDraw struct {
	DeckJson JsonDeck `json:"deckJson"`
	CardJson JsonCard `json:"cardJson"`
}
