package jsonschema

import (
	"testing"
)

func Test_ExtractJSONField(t *testing.T) {
	const simpleSchema = `{
	  "properties": {
		"firstName": {
		  "type": "string",
		  "minLength": 5
		},
		"lastName": {
		  "type": "string",
		  "maxLength": 7
		},
		"age": {
		  "type": "integer",
		  "minimum": 20,
		  "maximum": 45
		},
		"bankName": {
		  "enum": ["scb", "kbank"]
		},
		"bankAccountId": {
		  "type": "string",
		  "minLength": 11
		},
		"email": {
		  "format": "email"
		}
	  },
	  "required": ["firstName", "lastName", "age", "bankName", "bankAccountId"]
	}`

	const data = `{
		"firstName": "test1",
		"lastName": "test1",
		"age": 45,
		"bankName": "scb",
		"bankAccountId": "119534566125",
		"email": "test@gmail.com"
	}`

	t.Run("Happy - simple test get json field", func(t *testing.T) {
		Validate(simpleSchema, data)
	})

}
