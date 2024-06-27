# go-bruno-collection

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/speedscale/go-bruno-collection)

Go module to import and export Bruno collections.

This package provides a set of structures and import/export utilities for working with Bruno collections in Go. Use this package if you want to work natively with Bruno collection data and want basic data validation.

This project is not endorsed by or otherwise affiliated with Bruno project or company. For more information on the Bruno project, please visit their GitHub [repository](https://github.com/usebruno/bruno). We encourage you to support Bruno directly if desired. This project was inspired by, but shares no code with, its Postman equivalent [go-postman-collection](https://github.com/rbretecher/go-postman-collection/tree/master).

## Examples

### Collections

[Bruno](https://github.com/usebruno/bruno) imports and exports requests using a custom JSON format called a collection. These collections are well structured with data integrity rules for validation. This project only addresses working with exported Bruno collections in go and does not address working with Bruno's internal file structure. If you need to work with `.bru` files directly then you should go to the source as this project is not targeted at that use case.

#### Read a Bruno Collection

```go
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
```

#### Create and save a Bruno Collection

```go
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
```

## Contribution Guide

Feel free to open Merge Requests for any topic. This project is supported but not tightly controlled as far as features, bug fixes or enhancements. All contributions will be positively considered.