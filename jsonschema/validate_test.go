package jsonschema

import (
	"testing"
)

func Test_ExtractJSONField(t *testing.T) {
	const simpleSchema = `{
  "title": "Example Schema",
  "type": "object",
  "properties": {
    "firstName": {
      "type": "string"
    },
    "lastName": {
      "type": "string"
    },
    "age": {
      "description": "Age in years",
      "type": "integer",
      "minimum": 0
    }
  },
  "required": ["firstName", "lastName"]
}`

	const invalidPattern = `{
  "title": "Example Pattern",
  "type": "object",
  "properties": {
    "invalid": {
      "type": "string",
      "pattern": 99999
    }
  }
}`

	t.Run("Happy - simple test get json field", func(t *testing.T) {
		Validate(simpleSchema, invalidPattern)
	})

}
