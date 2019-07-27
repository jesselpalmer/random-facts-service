package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jesselpalmer/random-facts-api/routes"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	http.HandleFunc("/", routes.Facts)
	http.ListenAndServe(":"+port, nil)
}
