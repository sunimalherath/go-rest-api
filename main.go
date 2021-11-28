package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Coaster struct {
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	ID           string `json:"id"`
	InPark       string `json:"inPark"`
	Height       int    `json:"height"`
}

type coasterHandlers struct {
	store map[string]Coaster
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))

	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}

	jsonBytes, err := json.Marshal(coasters)
	if err != nil {

	}
	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Println(err)
	}
}

//
func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{
			"id1": {
				Name:         "Fury 324",
				Height:       100,
				ID:           "id1",
				InPark:       "LunaPark",
				Manufacturer: "BMW",
			},
		},
	}
}

func main() {
	cHandlers := newCoasterHandlers()
	http.HandleFunc("/coasters", cHandlers.get)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
