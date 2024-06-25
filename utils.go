package collection

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)

// Contains small helper functions to make collections slightly easier to work with. These
// functions are not necessary to use the core data structures.

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate performs data integrity checking based on the schema definition provided by the Bruno project.
func Validate(cs CollectionSchema) error {
	return validate.Struct(cs)
}

// Parse is a simple helper function that unmarshals the input bytes into a CollectionSchema
// whille also validating for data integrity. The caller can ignore validation errors if desired as the
// CollectionSchema is returned regardless.
func Parse(b []byte) (CollectionSchema, error) {
	cs := CollectionSchema{}
	if err := json.Unmarshal(b, &cs); err != nil {
		return cs, fmt.Errorf("failed to unmarshal CollectionSchema: %w", err)
	}

	return cs, validate.Struct(cs)
}

// ParseFile is a simple helper funtion that reads a file, unmarshals the results into a CollectionSchema
// and then performs basic data validation. The caller can ignore validation errors if desired as the
// CollectionSchema is returned regardless.
func ParseFile(filename string) (CollectionSchema, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return CollectionSchema{}, fmt.Errorf("failed to read file: %w", err)
	}

	return Parse(b)
}

// ToFile is a simple helper function that marshals the CollectionSchema and writes it to a file.
func WriteFile(filename string, cs CollectionSchema) error {
	b, err := json.Marshal(cs)
	if err != nil {
		return fmt.Errorf("failed to marshal CollectionSchema: %w", err)
	}

	if err = os.WriteFile(filename, b, 0644); err != nil {
		return fmt.Errorf("failed to write CollectionSchema to file: %w", err)
	}

	return nil
}
