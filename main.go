package main

import (
	"fmt"
	"net/http"

	"github.com/jesselpalmer/random-facts-api/routes"
)

func main() {
	http.HandleFunc("/", routes.Facts)
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
