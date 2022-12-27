package converter

import (
	"reflect"

	reflectutil "github.com/Bofry/structproto/util/reflectutil"
	"github.com/Bofry/types"
)

var (
	typeOfRawContent = reflect.TypeOf(types.RawContent(nil))
)

func RawContent(from interface{}) (types.RawContent, error) {
	if T, ok := from.(types.RawContent); ok {
		return T, nil
	} else if T, ok := from.(*types.RawContent); ok {
		return *T, nil
	} else if T, ok := from.([]byte); ok {
		return convBytesToRawContent(T)
	} else if T, ok := from.(string); ok {
		return convStringToRawContent(T)
	}

	rv := reflect.ValueOf(reflectutil.Indirect(from))
	switch rv.Kind() {
	case reflect.String:
		return convStringToRawContent(rv.String())
	case reflect.Array, reflect.Slice:
		if rv.CanInterface() {
			if T, ok := rv.Interface().([]byte); ok {
				return convBytesToRawContent(T)
			}
		}
	case reflect.Struct:
		if rv.Type().ConvertibleTo(typeOfUrl) {
			valueConv := rv.Convert(typeOfUrl)
			if valueConv.CanInterface() {
				return valueConv.Interface().(types.RawContent), nil
			}
		}
	}
	return nil, newConvErr(from, "types.RawContent")
}

func convStringToRawContent(value string) (types.RawContent, error) {
	return []byte(value), nil
}

func convBytesToRawContent(value []byte) (types.RawContent, error) {
	return value, nil
}
