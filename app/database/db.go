package database

import (
	"deckofcards/app/deck"
	"deckofcards/app/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	_ "github.com/lib/pq"
)

type PostDB interface {
	Open() error
	Close() error
	CreateDeck(deck *models.Deck) (int64, error)
	CreateCards(cards *models.Doc) error
	GetCardsByDeckId(deckId string) deck.Cards
	//GetDeckByDeckId(deckid string)
	DecrementRemainingCountOnDeckById(deckid string) error
}

type DB struct {
	db *sqlx.DB
}

// Opens connection with DB.
func (d *DB) Open() error {
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		log.Printf("Could not connect to Postgres DB...")
		return err
	}
	log.Println("Connected to the Database...")

	pg.MustExec(createSchema)

	d.db = pg

	return nil
}

// Closes the connection with DB.
func (d *DB) Close() error {
	return d.db.Close()
}

// Inserts New Deck.
func (d *DB) CreateDeck(deck *models.Deck) (int64, error) {
	lastId := 0
	err := d.db.QueryRow(insertNewDeck,
		deck.DeckId,
		deck.Shuffled,
		deck.Remaining,
		deck.CreatedDate,
		deck.LastModifiedDate).Scan(&lastId)
	if err != nil {
		log.Fatal("Deck Could not be created")
	}
	return int64(lastId), err
}

// Inserts Cards with deckId as primary key.
func (d *DB) CreateCards(cards *models.Doc) error {
	j, errr := json.Marshal(cards.DeckOfCards)
	if errr != nil {
		fmt.Printf("Error: %s", errr.Error())
	}
	_, err := d.db.Exec(insertCards, cards.DeckId, j)
	if err != nil {
		log.Fatal("Cards could not be created")
	}
	return err
}

func (d *DB) GetCardsByDeckId(deckId string) deck.Cards {

	var lang types.JSONText
	row := d.db.QueryRow(getCardsById, deckId)
	if err := row.Scan(&lang); err != nil {
		log.Fatalf("Row could not be selected: %v", err)
	}

	var v deck.Cards
	if err := json.Unmarshal(lang, &v); err != nil {
		log.Fatalf("Error on unmarshalling: %v", err)
	}

	return v
}

func (d *DB) DecrementRemainingCountOnDeckById(deckid string) error {
	_, err := d.db.Exec(decrementRemainingById, deckid)
	return err
}

/*
func (d *DB) GetDeckByDeckId(deckid string) models.Deck {

	var lang models.Deck
	row := d.db.QueryRow(getDeckById, deckid)

	if err := row.Scan(&lang.ID, &lang.DeckId, &lang.Shuffled,
		&lang.Remaining, &lang.CreatedDate, &lang.LastModifiedDate); err != nil {
		log.Fatalf("Row could not be selected: \n%v\n", err)
	}
	return lang
	//fmt.Printf("Deck Returned from db: \n%v\n", lang)
}
*/
