package oscar

import (
	"encoding/csv"
	"log"
	"os"
)

func ActorWhoGotMoreThanOne(filename string) map[string]int {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	cache := map[string]int{}

	for _, record := range data[1:] {
		cache[record[3]]++
	}

	for name, count := range cache {
		if count <= 1 {
			delete(cache, name)
		}
	}

	return cache
}
