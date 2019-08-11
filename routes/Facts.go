package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jesselpalmer/random-facts-service/models/fact"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
)

func getFactsList() []fact.Fact {
	fact1 := fact.Fact{
		Id:     0,
		Fact:   "The Unicorn is the national animal of Scotland",
		Source: "http://www.scotsman.com/heritage/people-places/scottish-fact-of-the-week-scotland-s-official-animal-the-unicorn-1-2564399",
	}
	fact2 := fact.Fact{
		Id:     1,
		Fact:   "In Switzerland it is illegal to own just one guinea pig",
		Source: "https://www.thefactsite.com/top-100-random-funny-facts/",
	}

	randomFactList := make([]fact.Fact, 0)
	randomFactList = append(randomFactList, fact1)
	randomFactList = append(randomFactList, fact2)

	return randomFactList
}

func getFactsLocally() []fact.Fact {
	log.Printf("getting facts locally")

	return getFactsList()
}

func getUnmarshaledFact(factItem *memcache.Item) []fact.Fact {
	log.Printf("memcache hit %s", factItem.Value)
	var fact []fact.Fact
	json.Unmarshal(factItem.Value, &fact)

	return fact
}

func cacheFacts(r *http.Request) []fact.Fact {
	log.Printf("caching facts")

	ctx := appengine.NewContext(r)

	bytes, err := json.Marshal(getFactsLocally())

	newFactItem := &memcache.Item{
		Key:   "facts",
		Value: []byte(bytes),
	}

	if err := memcache.Set(ctx, newFactItem); err != nil {
		log.Printf("error: could not cache facts")
		u := getFactsLocally()
		return u
	}

	log.Printf("facts cached")

	newCachedFactItem, err := memcache.Get(ctx, "facts")

	if err == nil {
		return getUnmarshaledFact(newCachedFactItem)
	}

	log.Printf("error: try again later")

	return []fact.Fact{}
}

func getFactsFromCache(r *http.Request) []fact.Fact {
	ctx := appengine.NewContext(r)

	factItem, err := memcache.Get(ctx, "facts")

	if err == nil {
		return getUnmarshaledFact(factItem)
	}

	log.Printf("error: memcache miss")
	u := cacheFacts(r)

	return u
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
		u := getFactsLocally()
		json.NewEncoder(w).Encode(u)
	}
}
