package utils

import (
	"fmt"
	"reflect"
)

func GetPropertyStruct(obj interface{}, fieldName string) (reflect.Value, error) {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		return reflect.Value{}, fmt.Errorf("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		return reflect.Value{}, fmt.Errorf("not found:" + fieldName)
	}
	return curField, nil
}
