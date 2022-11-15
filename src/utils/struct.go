package utils

import (
	"fmt"
	"reflect"
)

func GetPropertyStruct(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		fmt.Println("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		fmt.Println("not found:" + fieldName)
	}
	return curField
}
