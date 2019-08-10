package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jesselpalmer/random-facts-service/models/fact"
)

func getFacts() fact.Fact {
	u := fact.Fact{Data: "The Unicorn is the national animal of Scotland",
		Source: "http://www.scotsman.com/heritage/people-places/scottish-fact-of-the-week-scotland-s-official-animal-the-unicorn-1-2564399"}
	return u
}

// Facts : facts data route
func Facts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	u := getFacts()
	json.NewEncoder(w).Encode(u)
}
