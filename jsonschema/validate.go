package jsonschema

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"log"
)

func Validate(schema string, document string) {

	schemaLoader := gojsonschema.NewStringLoader(schema)
	documentLoader := gojsonschema.NewStringLoader(document)

	sl := gojsonschema.NewSchemaLoader()
	sl.Draft = gojsonschema.Draft7
	sl.AutoDetect = true
	sl.Validate = true
	sl.Compile(schemaLoader)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err.Error())
	}

	log.Printf("====> %+v", result)

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
