package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/jesselpalmer/random-apis/models/randomdata"
)

func index(w http.ResponseWriter, r *http.Request) {
	welcomeMessage := "Welcome to the Random APIs!"
	fmt.Fprintf(w, "%s", welcomeMessage)
}

func facts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Random Fact Generator")
}

func factsData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	u := randomdata.RandomData{Data: "The Unicorn is the national animal of Scotland",
		Resource: "http://www.scotsman.com/heritage/people-places/scottish-fact-of-the-week-scotland-s-official-animal-the-unicorn-1-2564399",
		Type:     "fact"}

	json.NewEncoder(w).Encode(u)
}

func thoughts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Random Thought Generator")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/facts", facts)
	http.HandleFunc("/facts/", factsData)
	http.HandleFunc("/thoughts", thoughts)
	http.ListenAndServe(":8080", nil)
}
