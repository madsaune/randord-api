package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type application struct {
	infoLog  log.Logger
	errorLog log.Logger
	wordlist []string
}

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", app.indexHandler)
	router.HandleFunc("/{count}", app.indexHandler)

	return router
}

func (app *application) indexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	countVar := vars["count"]

	// if count is not set or not a number
	// we default to 10 words
	count, err := strconv.Atoi(countVar)
	if err != nil {
		count = 10
	}

	// Limit result to 100 words
	if count > 100 {
		count = 100
	}

	// Get the prefered format
	acceptHeader := r.Header.Get("Accept")

	// Generate list of words
	var words []string
	for i := 0; i < count; i += 1 {
		randIdx := randInt(0, len(app.wordlist)-1)
		word := app.wordlist[randIdx]
		words = append(words, strings.Title(word))
	}

	resp := response{
		Count:  count,
		Words:  words,
		Format: acceptHeader,
	}

	resp.WriteResponse(w)
}
