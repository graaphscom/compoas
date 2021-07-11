package compoas

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDump(t *testing.T) {
	testCases := []struct {
		prettyPrint bool
		expected    string
		name        string
	}{
		{
			name:        "Pretty print",
			prettyPrint: true,
			expected:    "{\n  \"openapi\": \"\",\n  \"info\": {\n    \"title\": \"\",\n    \"version\": \"\"\n  },\n  \"paths\": null\n}",
		},
		{
			name:        "No pretty print",
			prettyPrint: false,
			expected:    "{\"openapi\":\"\",\"info\":{\"title\":\"\",\"version\":\"\"},\"paths\":null}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := OAS{}.Dump(testCase.prettyPrint, "test_dump.json")
			assert.NoError(t, err)
			dump, err := os.ReadFile("test_dump.json")
			assert.NoError(t, err)
			assert.Equal(t, testCase.expected, string(dump))
		})
	}
	os.Remove("test_dump.json")
}

func TestMerge(t *testing.T) {
	result := (&OAS{
		Components: &Components{
			Schemas:         map[string]Schema{"initialSchema": {}},
			SecuritySchemes: map[string]SecurityScheme{},
		},
		Paths: map[string]PathItem{"initialPath": {}},
	}).Merge(OAS{
		Components: &Components{
			Schemas:         map[string]Schema{"mergedSchema": {}},
			SecuritySchemes: map[string]SecurityScheme{"bearerAuth": {}},
		},
		Paths: map[string]PathItem{"mergedPath": {}},
	})
	assert.Equal(
		t,
		&OAS{
			Components: &Components{
				Schemas:         map[string]Schema{"initialSchema": {}, "mergedSchema": {}},
				SecuritySchemes: map[string]SecurityScheme{"bearerAuth": {}},
			},
			Paths: map[string]PathItem{"initialPath": {}, "mergedPath": {}},
		},
		result,
	)
}
