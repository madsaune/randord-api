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

	var c int
	var err error

	count, ok := vars["count"]
	if !ok {
		c = 10
	} else {
		c, err = strconv.Atoi(count)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}

	acceptHeader := r.Header.Get("Accept")
	if acceptHeader == "" {
		acceptHeader = "application/json"
	}

	var words []string

	for i := 0; i < c; i += 1 {
		word := app.wordlist[randInt(0, len(app.wordlist)-1)]
		words = append(words, strings.Title(word))
	}

	resp := response{
		Count: c,
		Words: words,
	}

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
			app.errorLog.Fatalf("could not build json response: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
