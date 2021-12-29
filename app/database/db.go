package database

import (
	"deckofcards/app/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostDB interface {
	Open() error
	Close() error
	CreateDeck(deck *models.Deck) (int64, error)
	CreateCards(cards *models.Doc) error
	GetByDeckId(deckId string)
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

func (d *DB) GetByDeckId(deckId string) {
	res, err := d.db.Query(getDeckById, deckId)
	if err != nil {
		log.Fatal("Could not retrive deck by id")
	}
	fmt.Println(res)
}
