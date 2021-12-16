package database

const createSchema = `
	CREATE TABLE IF NOT EXISTS decks
	(
		id SERIAL PRIMARY KEY,
		deckId SERIAL,
		belongsTo TEXT,
		createdDate DATE,
		sessionId TEXT
	)
`
