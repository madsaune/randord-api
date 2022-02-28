package main

import (
	"encoding/json"
	"strings"
)

type response struct {
	Count int      `json:"count"`
	Words []string `json:"words"`
}

func (r *response) EncodeJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *response) EncodePlain() []byte {
	return []byte(strings.Join(r.Words, ", "))
}
