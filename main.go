package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jesselpalmer/random-facts-service/routes"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/status", statusHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	routes.Facts(w, r)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
