package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type T struct {
	Name  string
	Value *json.RawMessage
}

var stringType = reflect.TypeOf("")
var jsonRawMessageType = reflect.TypeOf((*json.RawMessage)(nil))

//func TestIdiot() (*string, error) {
//	tagData := string(`{"name": "idiot=name,required,min=1,max=5"}`)
//
//	args := strings.Split(string(tagData), ":")
//	log.Printf("===> %+v", args[0])
//
//	dd := d{"title"}
//
//	idiot.Validator.Validate(&dd)
//
//	result := args[0]
//	return &result, nil
//}

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil

}

//func isJSON(s string) bool {
//	var js map[string]interface{}
//	return json.Unmarshal([]byte(s), &js) == nil
//
//}

func isJSONRaw(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func main() {
	//var tests = []string{
	//	`"Platypus"`,
	//	`Platypus`,
	//	`{"id":"idiot:id,required,min=1,max=5"}`,
	//}
	//
	//for _, t := range tests {
	//	fmt.Printf("isJSONString(%s) = %v\n", t, isJSONString(t))
	//	fmt.Printf("isJSON(%s) = %v\n\n", t, isJSON(t))
	//	log.Printf("===> isJSONRaw %+v", isJSONRaw(t))
	//}

	j := string(`{"name":"idiot=required,min=1,max=5"}`)

	if isJSON(j) {
		log.Println("====> error not json")
	}

	//t := T{Name: "name", Value: j}
	sp := strings.Split(j, ":")

	log.Printf("====> raw %+v", j)
	log.Printf("====> raw split %+v", sp[1])

	for i := 0; i < len(sp); i++ {
		log.Printf("====> i %+v", sp[i])
	}

	//v := reflect.Indirect(reflect.ValueOf(&t))
	//v = v.FieldByName("Value")

	//test(v)

}

func test(v reflect.Value) {
	switch v.Type() {
	case stringType:
		fmt.Println("String")
	case jsonRawMessageType:
		fmt.Printf("*json.RawMessage: v: %+v , v.Kind(): %+v, v.Type(): %+v\n", v, v.Kind(), v.Type())
	default:
		fmt.Printf("Default: v: %+v , v.Type(): %+v\n", v, v.Type())
	}

}

func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}
