package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/jesselpalmer/random-facts-api/routes"
)


func main() {
	http.HandleFunc("/", routes.Facts)
	http.HandleFunc("/_ah/health", healthCheckHandler)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
