package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "server/api/models"
	"server/utils"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

func main() {
	// initialize router
	r := mux.NewRouter()

	// router endpoints
	r.HandleFunc("/", getUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))

}

type Result struct {
	Results []User	`json:"results"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	payload := &Result{}

	resp, err := client.Get("https://randomuser.me/api/?results=1")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := json.NewDecoder(resp.Body).Decode(payload); err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")

	utils.Respond(w, payload.Results)
}
