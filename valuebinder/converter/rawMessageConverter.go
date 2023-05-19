package converter

import (
	"encoding/json"
	"reflect"
)

var (
	typeOfRawMessage = reflect.TypeOf(json.RawMessage(nil))
)

func RawMessage(from interface{}) (json.RawMessage, error) {
	if T, ok := from.(json.RawMessage); ok {
		return T, nil
	} else if T, ok := from.(*json.RawMessage); ok {
		return *T, nil
	} else if T, ok := from.([]byte); ok {
		return convBytesToRawMessage(T)
	} else if T, ok := from.(string); ok {
		return convStringToRawMessage(T)
	}

	rv := reflect.ValueOf(indirect(from))
	switch rv.Kind() {
	case reflect.String:
		return convStringToRawMessage(rv.String())
	case reflect.Array, reflect.Slice:
		if rv.CanInterface() {
			if T, ok := rv.Interface().([]byte); ok {
				return convBytesToRawMessage(T)
			}
		}
	}
	return nil, newConvErr(from, "json.RawMessage")
}

func convStringToRawMessage(value string) (json.RawMessage, error) {
	return []byte(value), nil
}

func convBytesToRawMessage(value []byte) (json.RawMessage, error) {
	return value, nil
}
