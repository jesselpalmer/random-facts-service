package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jesselpalmer/random-facts-service/models/fact"
)

// Facts : facts data route
func Facts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	u := fact.Fact{Data: "The Unicorn is the national animal of Scotland",
		Source: "http://www.scotsman.com/heritage/people-places/scottish-fact-of-the-week-scotland-s-official-animal-the-unicorn-1-2564399"}

	json.NewEncoder(w).Encode(u)
}
