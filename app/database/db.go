package database

import (
	"deckofcards/app/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostDB interface {
	Open() error
	Close() error
	CreateDeck(deck *models.Deck) (int64, error)
	CreateCards(cards *models.Cards) error
}

type DB struct {
	db *sqlx.DB
}

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

func (d *DB) Close() error {
	return d.db.Close()
}

// Creates
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

func (d *DB) CreateCards(cards *models.Cards) error {

	_, err := d.db.Exec(insertCards, cards.DeckId, cards.DeckOfCards)
	if err != nil {
		log.Fatal("Cards could not be created")
	}
	return err
}
