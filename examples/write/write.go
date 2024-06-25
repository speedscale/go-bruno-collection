package main

import (
	"os"

	collection "github.com/speedscale/go-bruno-collection"
)

func main() {
	cs := collection.CreateCollection("demo", "Demo collection")
	collection.AddItem(&cs, collection.ItemSchema{
		Type:    "http",
		Name:    "test",
		Request: collection.CreateRequest("http://example.com", "GET"),
	})

	err := collection.WriteFile("test_bruno_collection.json", cs)
	if err != nil {
		panic(err)
	}
	defer os.Remove("test_bruno_collection.json")
}
