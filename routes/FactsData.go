package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jesselpalmer/random-apis/models/randomdata"
)

// FactsData : facts data route
func FactsData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	u := randomdata.RandomData{Data: "The Unicorn is the national animal of Scotland",
		Resource: "http://www.scotsman.com/heritage/people-places/scottish-fact-of-the-week-scotland-s-official-animal-the-unicorn-1-2564399",
		Type:     "fact"}

	json.NewEncoder(w).Encode(u)
}
