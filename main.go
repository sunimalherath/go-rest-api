package main

import (
	"encoding/json"
	"io/ioutil"
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

func (h *coasterHandlers) coasters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application-json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonBytes)
	if err != nil {
		log.Println(err)
	}
}

func (h *coasterHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var coaster Coaster
	err = json.Unmarshal(bodyByte, &coaster)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	//h.Lock()
	//defer h.Unlock()
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
	http.HandleFunc("/coasters", cHandlers.coasters)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
