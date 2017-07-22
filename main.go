package main

import (
	"fmt"
	"net/http"

	"github.com/jesselpalmer/routes"
)

func index(w http.ResponseWriter, r *http.Request) {
	welcomeMessage := "Welcome to the Random APIs!"
	fmt.Fprintf(w, "%s", welcomeMessage)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/facts", routes.Facts)
	http.HandleFunc("/facts/", routes.FactsData)
	http.HandleFunc("/greetings", routes.Greetings)
	http.HandleFunc("/thoughts", routes.Thoughts)
	http.ListenAndServe(":8080", nil)
}
