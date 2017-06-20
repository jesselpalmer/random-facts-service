package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var welcomeMessage = "Welcome to the Random APIs!"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", welcomeMessage)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
