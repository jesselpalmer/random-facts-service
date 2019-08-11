package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jesselpalmer/random-facts-service/models/fact"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

var randomFact = fact.Fact{
	Id:     0,
	Fact:   "The Unicorn is the national animal of Scotland",
	Source: "http://www.scotsman.com/heritage/people-places/scottish-fact-of-the-week-scotland-s-official-animal-the-unicorn-1-2564399",
}

func cacheFacts(r *http.Request) fact.Fact {
	log.Printf("caching facts")

	ctx := appengine.NewContext(r)

	bytes, err := json.Marshal(randomFact)

	item1 := &memcache.Item{
		Key:   "facts",
		Value: []byte(bytes),
	}

	if err := memcache.Set(ctx, item1); err != nil {
		log.Printf("error: could not cache facts")
		u := getFactsLocally(r)
		return u
	}

	log.Printf("facts cached")

	item0, err := memcache.Get(ctx, "facts")

	if err == nil {
		log.Printf("memcache hit %s", item0.Value)
		var fact fact.Fact
		json.Unmarshal(item0.Value, &fact)
		return fact
	}

	log.Printf("error: try again later")
	return fact.Fact{}
}

func getFactsFromCache(r *http.Request) fact.Fact {
	ctx := appengine.NewContext(r)

	item0, err := memcache.Get(ctx, "facts")

	if err == nil {
		log.Printf("memcache hit %s", item0.Value)
		var fact fact.Fact
		json.Unmarshal(item0.Value, &fact)
		return fact
	}

	log.Printf("error: memcache miss")
	u := cacheFacts(r)

	return u
}

func getFactsLocally(r *http.Request) fact.Fact {
	log.Printf("getting facts locally")
	return randomFact
}

// Facts : facts data route
func Facts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if appengine.IsAppEngine() {
		u := getFactsFromCache(r)
		json.NewEncoder(w).Encode(u)
	} else {
		log.Printf("app not in GAE environment")
		u := getFactsLocally(r)
		json.NewEncoder(w).Encode(u)
	}
}
