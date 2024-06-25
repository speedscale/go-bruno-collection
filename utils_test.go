package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// XXX: These unit tests are simple sanity checks for small utility functions, they are not
// (currently) intended to perform comprehensive data validation.

func TestParse(t *testing.T) {
	validJSON := []byte(`{"name":"valid schema", "version":"1"}`)
	invalidJSON := []byte(`{"unknown key":""}`)
	badJSON := []byte(`{"name":`)

	tests := []struct {
		name     string
		input    []byte
		expected CollectionSchema
		hasError bool
	}{
		{
			name:     "Valid JSON",
			input:    validJSON,
			expected: CollectionSchema{Name: "valid schema", Version: "1"},
			hasError: false,
		},
		{
			name:     "Invalid JSON",
			input:    invalidJSON,
			expected: CollectionSchema{Name: ""},
			hasError: true,
		},
		{
			name:     "Bad JSON",
			input:    badJSON,
			expected: CollectionSchema{},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs, err := Parse(tt.input)
			assert.Equal(t, tt.expected, cs)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseFile(t *testing.T) {
	validJSON := `{"name":"valid schema"}`
	invalidJSON := `{"name":""}`

	tests := []struct {
		name              string
		filename          string
		content           string
		expectedItemCount int
		expectError       bool
	}{
		{
			name:              "Valid File",
			filename:          "./testdata/demo.json",
			content:           validJSON,
			expectedItemCount: 8,
			expectError:       false,
		},
		{
			name:        "Invalid File",
			filename:    "invalid.json",
			content:     invalidJSON,
			expectError: true,
		},
		{
			name:        "Non-existent File",
			filename:    "nonexistent.json",
			content:     "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with the test content if it is not supposed to be non-existent.
			cs, err := ParseFile(tt.filename)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedItemCount, cs.Items)
		})
	}
}
