package routes

import (
	"fmt"
	"net/http"
)

// Thoughts : thoughts data route
func Thoughts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Random Thought Generator")
}
