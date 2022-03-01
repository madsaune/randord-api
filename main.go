package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	// -----------------------------------------------------------------------------
	//     - CONFIG -
	// -----------------------------------------------------------------------------

	rand.Seed(time.Now().UTC().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// -----------------------------------------------------------------------------
	//     - LOGGERS -
	// -----------------------------------------------------------------------------

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// -----------------------------------------------------------------------------
	//     - WORDLIST -
	// -----------------------------------------------------------------------------

	wl, err := readWordlist("./data/wordlist.txt")
	if err != nil {
		errorLog.Fatalf("could not read wordlist: %v", err)
	}

	// -----------------------------------------------------------------------------
	//     - INIT APPLICATION -
	// -----------------------------------------------------------------------------

	app := &application{
		infoLog:  *infoLog,
		errorLog: *errorLog,
		wordlist: wl,
	}

	srv := &http.Server{
		Addr:     ":" + port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on localhost:%s ...\n", port)
	errorLog.Fatal(srv.ListenAndServe())
}
