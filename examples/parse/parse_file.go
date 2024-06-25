package main

import (
	collection "github.com/speedscale/go-bruno-collection"
)

func main() {
	cs, err := collection.ParseFile("../testdata/demo.json")
	if err != nil {
		panic(err)
	}

	_ = cs
}
