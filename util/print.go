package util

import (
	"fmt"
	"reflect"
)

func PrintStruct(secret interface{}) {
	value := reflect.ValueOf(secret)
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}
}
