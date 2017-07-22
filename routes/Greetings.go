package routes

import (
	"fmt"
	"net/http"
)

// Greetings : greetings data route
func Greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Random Greeting Generator")
}
