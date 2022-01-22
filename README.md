# DeckOfCards API

Simple API for cards games. Creates a Deck of Cards consisting 52 playing cards. Newly created deck is not shuffled. It is an ordered deck.

- Creates brand new deck of cards, assigns it to an DeckID.

  ```
  POST /api/createDeck
  ```

- Shuffles the cards by appointed DeckID.
  ```
  POST /api/shuffle/{deckid}
  ```
- Draw a single card from the deck via DeckID.
  ```
  GET /api/draw/{deckid}
  ```
  <br>

## For starting Development environment:

**_!! Important note !!:_**

- Docker must be installed and PORTS 5050, 5432 and 8080 must be opened.

  ```
  Start:

  $ docker-compose up -d --build


  Stop:

  $ docker-compose down
  ```

- This will create 3 containers:

  postgres_container: PostgresSQL db

  pgadmin_container: PG Admin

  deckofcards-go: API

<br>

Thanks, have fun :)
