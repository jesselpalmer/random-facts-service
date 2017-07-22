package routes

import (
	"fmt"
	"net/http"
)

// Facts : facts route
func Facts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "Random Fact Generator")
}
