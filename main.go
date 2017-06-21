package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	var welcomeMessage = "Welcome to the Random APIs!"
	fmt.Fprintf(w, "%s", welcomeMessage)
}

func facts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Random Fact Generator")
}

func factsData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Random data is returned here")
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
