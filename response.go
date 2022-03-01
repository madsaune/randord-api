package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type response struct {
	Count  int      `json:"count"`
	Words  []string `json:"words"`
	Format string   `json:"-"`
}

func (r *response) WriteResponse(w http.ResponseWriter) {
	if r.Format == "text/plain" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(strings.Join(r.Words, "\n")))
		return
	}

	// TODO: replace with html
	if r.Format == "text/html" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(strings.Join(r.Words, ", ")))
		return
	}

	// defaults to JSON
	data, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
