package main

import (
	"reflect"
)

// 节点类型
func nodeType(value interface{}) NodeType {
	if reflect.TypeOf(value).Kind() == reflect.Int ||
		reflect.TypeOf(value).Kind() == reflect.Int8 ||
		reflect.TypeOf(value).Kind() == reflect.Uint8 ||
		reflect.TypeOf(value).Kind() == reflect.Int16 ||
		reflect.TypeOf(value).Kind() == reflect.Uint16 ||
		reflect.TypeOf(value).Kind() == reflect.Int32 ||
		reflect.TypeOf(value).Kind() == reflect.Uint32 ||
		reflect.TypeOf(value).Kind() == reflect.Int64 ||
		reflect.TypeOf(value).Kind() == reflect.Uint64 {
		return TypeInt
	}

	if reflect.TypeOf(value).Kind() == reflect.Int64 ||
		reflect.TypeOf(value).Kind() == reflect.Uint64 {
		return TypeDouble
	}

	if reflect.TypeOf(value).Kind() == reflect.String {
		return TypeString
	}

	if reflect.TypeOf(value).Kind() == reflect.Bool {
		return TypeBool
	}

	if reflect.TypeOf(value).Kind() == reflect.Array {
		return TypeArray
	}

	if reflect.TypeOf(value).Kind() == reflect.Float32 {
		return TypeFloat
	}

	if reflect.TypeOf(value).Kind() == reflect.Float64 {
		return TypeDouble
	}

	return TypeNone
}
