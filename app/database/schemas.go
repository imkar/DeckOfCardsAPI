package database

const createSchema = `
	CREATE TABLE IF NOT EXISTS decks
	(
		id SERIAL,
		deckId VARCHAR(7) PRIMARY KEY UNIQUE,
		shuffled BOOLEAN,
		remaining INTEGER,
		createdDate DATE,
		lastModifiedDate DATE
	);
	CREATE TABLE IF NOT EXISTS cards
	(
		deckId VARCHAR(7) ,
		deckOfCards json,
		CONSTRAINT fk_decks
		FOREIGN KEY(deckId) 
		REFERENCES decks(deckId)
		ON DELETE SET NULL
	)
`

var insertNewDeck = `
	INSERT INTO decks(deckId, shuffled, remaining, createdDate, lastModifiedDate) VALUES($1,$2,$3,$4,$5) RETURNING id
`

var insertCards = `
	INSERT INTO cards (deckId, deckOfCards) VALUES ($1, $2)
`
