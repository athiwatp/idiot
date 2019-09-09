package idiot

import (
	json2 "encoding/json"
	"github.com/tidwall/gjson"
	"log"
	"testing"
)

func Test_ExtractJSONField(t *testing.T) {
	json := json2.RawMessage(`{"name":{"first":"required,min=1,max=5,enum=KBANK,SCB","last":"Prichard"},"age":47}`)
	jsonNested := json2.RawMessage(`{
	  "programmers": [
		{
		  "firstName": "Janet", 
		  "lastName": "McLaughlin", 
		}, {
		  "firstName": "Elliotte", 
		  "lastName": "Hunter", 
		}, {
		  "firstName": "Jason", 
		  "lastName": "Harold", 
		}
	  ]
	}`)

	t.Run("Happy - simple test get json field", func(t *testing.T) {

		value := gjson.Get(string(json), "name.first")
		log.Printf("===> %+v", value)

		//sections := strings.Split(logMessage, "\n")
		//assert.Equal(t, "Query: var1=1&var2=2", sections[1])
		//assert.Equal(t, `Body: {"a":"[HiddenField]","b":2,"c":"3"}`, sections[2])
		//assert.Equal(t, "x-velo-auth: tokenString", sections[3])
	})

	t.Run("Happy - Nested json field", func(t *testing.T) {

		result := gjson.Get(string(jsonNested), "programmers.#.lastName")
		for _, name := range result.Array() {
			println(name.String())
		}
	})
}
