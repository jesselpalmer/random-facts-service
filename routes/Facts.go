package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func cacheFacts(r *http.Request) string {
	log.Printf("caching facts")

	ctx := appengine.NewContext(r)

	item1 := &memcache.Item{
		Key:   "facts",
		Value: []byte("The Unicorn is the national animal of Scotland"),
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
		u := string(item0.Value)
		return u
	}

	return "error: try again later"
}

func getFactsFromCache(r *http.Request) string {
	ctx := appengine.NewContext(r)

	item0, err := memcache.Get(ctx, "facts")

	if err == nil {
		log.Printf("memcache hit %s", item0.Value)
		u := string(item0.Value)
		return u
	}

	log.Printf("error: memcache miss")
	u := cacheFacts(r)

	return u
}

func getFactsLocally(r *http.Request) string {
	log.Printf("getting facts locally")
	return "The Unicorn is the national animal of Scotland"
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
