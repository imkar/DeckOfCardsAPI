package models

type Deck struct {
	ID               int64  `db:"id"`
	DeckId           string `db:"deckId"`
	Shuffled         bool   `db:"shuffled"`
	Remaining        uint8  `db:"remaining"`
	CreatedDate      string `db:"createdDate"`
	LastModifiedDate string `db:"lastModifiedDate"`
}

type JsonDeck struct {
	ID               int64  `json:"id"`
	DeckId           string `json:"deckId"`
	Shuffled         bool   `json:"shuffled"`
	Remaining        uint8  `json:"remaining"`
	CreatedDate      string `json:"createdDate"`
	LastModifiedDate string `json:"lastModifiedDate"`
}

type Cards struct {
	DeckId      string `json:"deckId"`
	DeckOfCards string `json:"deckOfCards"`
}
