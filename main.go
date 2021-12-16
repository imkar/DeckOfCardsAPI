package main

import (
	"deckofcards/app"
	"deckofcards/app/database"
	"log"
	"net/http"
	"os"
)

func main() {
	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err)

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("Deck Of Cards App running...")
	err = http.ListenAndServe(":9000", nil)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
