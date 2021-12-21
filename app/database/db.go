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
	CreateDeck(deck *models.Deck) error
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

func (d *DB) CreateDeck(deck *models.Deck) error {
	res, err := d.db.Exec(insertNewDeck)
	if err != nil {
		log.Fatal("Deck Could not be created")
	}
	res.LastInsertId()
	return err
}
