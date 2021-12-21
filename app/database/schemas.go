package database

const createSchema = `
	CREATE TABLE IF NOT EXISTS decks
	(
		id SERIAL,
		deckId VARCHAR(7) PRIMARY KEY,
		shuffled BOOLEAN,
		remaining INTEGER,
		createdDate DATE,
		lastModifiedDate DATE
	)
`

var insertNewDeck = `
	INSERT INTO decks(deckId, shuffled, remaining, createdDate, lastModifiedDate) VALUES($1,$2,$3,$4,$5) RETURNING id
`
