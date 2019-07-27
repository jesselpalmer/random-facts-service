package main

import (
	"fmt"
	"net/http"

	"github.com/jesselpalmer/random-apis/routes"
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
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
