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
	count, err := strconv.Atoi(countVar)
	if err != nil {
		// defaults to 10 words
		count = 10
	}

	// Limit result to 100 words
	if count > 100 {
		count = 100
	}

	// Get the prefered format
	acceptHeader := r.Header.Get("Accept")
	if acceptHeader == "" {
		acceptHeader = "application/json"
	}

	// Generate list of words
	var words []string
	for i := 0; i < count; i += 1 {
		word := app.wordlist[randInt(0, len(app.wordlist)-1)]
		words = append(words, strings.Title(word))
	}

	resp := response{
		Count: count,
		Words: words,
	}

	// Respond based on Accept header
	switch acceptHeader {
	case "text/html":
		w.Header().Set("Content-Type", "text/plain")
		w.Write(resp.EncodePlain())
	case "text/plain":
		w.Header().Set("Content-Type", "text/plain")
		w.Write(resp.EncodePlain())
	default:
		data, err := resp.EncodeJSON()
		if err != nil {
			app.errorLog.Printf("could not build json response: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
