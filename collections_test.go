package collection

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

// XXX (MEL): These unit tests address basic validation rules but are not comprehensive. An
// industrious person could definitely make this more complete.
func init() {
	validate = validator.New()
}

func TestCollectionSchemaValidation(t *testing.T) {
	tests := []struct {
		name        string
		collection  CollectionSchema
		expectError bool
	}{
		{
			name: "Valid CollectionSchema",
			collection: CollectionSchema{
				Version: "1",
				UID:     "uid", // Assume valid UIDSchema
				Name:    "Valid Name",
				Items: []ItemSchema{
					{
						UID:  "uid", // Assume valid UIDSchema
						Type: "http",
						Name: "Valid Item",
					},
				},
				ActiveEnvironmentUID: nil,
				Environments: []EnvironmentSchema{
					{
						UID:  "uid", // Assume valid UIDSchema
						Name: "Valid Environment",
						Variables: []EnvironmentVariablesSchema{
							{
								UID:     "uid", // Assume valid UIDSchema
								Type:    "text",
								Enabled: true,
							},
						},
					},
				},
				Pathname:            nil,
				RunnerResult:        nil,
				CollectionVariables: nil,
				BrunoConfig:         nil,
			},
			expectError: false,
		},
		{
			name: "Invalid Version",
			collection: CollectionSchema{
				Version: "2",   // Invalid version
				UID:     "uid", // Assume valid UIDSchema
				Name:    "Valid Name",
			},
			expectError: true,
		},
		{
			name: "Empty Name",
			collection: CollectionSchema{
				Version: "1",
				UID:     "uid", // Assume valid UIDSchema
				Name:    "",
			},
			expectError: true,
		},
		{
			name: "Invalid ActiveEnvironmentUID",
			collection: CollectionSchema{
				Version:              "1",
				UID:                  "uid", // Assume valid UIDSchema
				Name:                 "Valid Name",
				ActiveEnvironmentUID: ptrString("1234567890123456789"), // Invalid length
			},
			expectError: true,
		},
		{
			name: "Invalid Environment",
			collection: CollectionSchema{
				Version: "1",
				UID:     "uid", // Assume valid UIDSchema
				Name:    "Valid Name",
				Environments: []EnvironmentSchema{
					{
						UID:  "uid", // Assume valid UIDSchema
						Name: "",    // Invalid Name
					},
				},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.collection)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestEnvironmentVariablesSchemaJSON(t *testing.T) {
	original := EnvironmentVariablesSchema{
		UID:     "123",
		Name:    ptrString("test-name"),
		Value:   ptrString("test-value"),
		Type:    "text",
		Enabled: true,
		Secret:  ptrBool(true),
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(original)
	assert.NoError(t, err)

	// Unmarshal back to struct
	var unmarshaled EnvironmentVariablesSchema
	err = json.Unmarshal(jsonData, &unmarshaled)
	assert.NoError(t, err)

	// Check if original and unmarshaled are equal
	assert.Equal(t, original, unmarshaled)
}

func TestEnvironmentSchemaJSON(t *testing.T) {
	original := EnvironmentSchema{
		UID:  "123",
		Name: "env-name",
		Variables: []EnvironmentVariablesSchema{
			{
				UID:     "456",
				Name:    ptrString("var-name"),
				Value:   ptrString("var-value"),
				Type:    "text",
				Enabled: true,
				Secret:  ptrBool(false),
			},
		},
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(original)
	assert.NoError(t, err)

	// Unmarshal back to struct
	var unmarshaled EnvironmentSchema
	err = json.Unmarshal(jsonData, &unmarshaled)
	assert.NoError(t, err)

	// Check if original and unmarshaled are equal
	assert.Equal(t, original, unmarshaled)
}

func TestCollectionSchemaJSON(t *testing.T) {
	original := CollectionSchema{
		Version:              "1",
		UID:                  "123",
		Name:                 "collection-name",
		Items:                []ItemSchema{},
		ActiveEnvironmentUID: ptrString("env-uid"),
		Environments:         []EnvironmentSchema{},
		Pathname:             ptrString("/path/to/collection"),
		RunnerResult:         map[string]interface{}{"key": "value"},
		CollectionVariables:  map[string]interface{}{"var": "val"},
		BrunoConfig:          map[string]interface{}{"config": "conf"},
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(original)
	assert.NoError(t, err)

	// Unmarshal back to struct
	var unmarshaled CollectionSchema
	err = json.Unmarshal(jsonData, &unmarshaled)
	assert.NoError(t, err)

	// Check if original and unmarshaled are equal
	assert.Equal(t, original, unmarshaled)
}

func TestCreateRequest(t *testing.T) {
	t.Run("Valid URL and Method", func(t *testing.T) {
		req := CreateRequest("http://example.com", "GET")
		assert.Equal(t, "http://example.com", req.URL)
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "inherit", req.Auth.Mode)
		assert.Equal(t, "none", req.Body.Mode)
		// check the built in data validation as well (even though in theory it is redundant)
		assert.NoError(t, validate.Struct(req))
	})

	t.Run("Empty URL", func(t *testing.T) {
		req := CreateRequest("", "POST")
		assert.Equal(t, "", req.URL)
		assert.Equal(t, "POST", req.Method)
		assert.Equal(t, "inherit", req.Auth.Mode)
		assert.Equal(t, "none", req.Body.Mode)
		// check the built in data validation as well (even though in theory it is redundant)
		assert.Error(t, validate.Struct(req))
	})

	t.Run("Empty Method", func(t *testing.T) {
		req := CreateRequest("http://example.com", "")
		assert.Equal(t, "http://example.com", req.URL)
		assert.Equal(t, "", req.Method)
		assert.Equal(t, "inherit", req.Auth.Mode)
		assert.Equal(t, "none", req.Body.Mode)
		// check the built in data validation as well (even though in theory it is redundant)
		assert.Error(t, validate.Struct(req))
	})
}

func ptrString(s string) *string {
	return &s
}

func ptrBool(b bool) *bool {
	return &b
}
