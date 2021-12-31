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

var getCardsById = `
	SELECT deckofcards FROM decks JOIN cards ON (decks.deckId=cards.deckId) WHERE decks.deckid = $1;
`

var getDeckById = `
	SELECT * FROM decks WHERE deckid = $1;
`

var decrementRemainingById = `
	UPDATE decks SET remaining = remaining - 1 WHERE deckid = $1;
`

var updateCardsById = `
	UPDATE cards SET deckofcards = $1 WHERE deckid = $2;
`

var deleteCardsById = `
	DELETE FROM cards WHERE deckid = $1;
`

var deleteDeckById = `
	DELETE FROM decks WHERE deckid = $1;
`
