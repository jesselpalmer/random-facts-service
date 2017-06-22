package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RandomData struct {
	data     string
	resource string
}

func index(w http.ResponseWriter, r *http.Request) {
	var welcomeMessage = "Welcome to the Random APIs!"
	fmt.Fprintf(w, "%s", welcomeMessage)
}

func facts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Random Fact Generator")
}

func factsData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	u := RandomData{data: "The Unicorn is the national animal of Scotland",
		resource: "http://www.scotsman.com/heritage/people-places/scottish-fact-of-the-week-scotland-s-official-animal-the-unicorn-1-2564399"}

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
