package converter

import (
	"bytes"
	"reflect"

	reflectutil "github.com/Bofry/structproto/util/reflectutil"
)

var (
	typeOfBuffer = reflect.TypeOf(bytes.Buffer{})
)

func Buffer(from interface{}) (bytes.Buffer, error) {
	if T, ok := from.(bytes.Buffer); ok {
		return T, nil
	} else if T, ok := from.(*bytes.Buffer); ok {
		return *T, nil
	} else if T, ok := from.([]byte); ok {
		return convBytesToBuffer(T)
	} else if T, ok := from.(string); ok {
		return convStringToBuffer(T)
	}

	rv := reflect.ValueOf(reflectutil.Indirect(from))
	switch rv.Kind() {
	case reflect.String:
		return convStringToBuffer(rv.String())
	case reflect.Array, reflect.Slice:
		if rv.CanInterface() {
			if T, ok := rv.Interface().([]byte); ok {
				return convBytesToBuffer(T)
			}
		}
	case reflect.Struct:
		if rv.Type().ConvertibleTo(typeOfBuffer) {
			valueConv := rv.Convert(typeOfBuffer)
			if valueConv.CanInterface() {
				return valueConv.Interface().(bytes.Buffer), nil
			}
		}
	}
	return bytes.Buffer{}, newConvErr(from, "bytes.Buffer")
}

func convStringToBuffer(value string) (bytes.Buffer, error) {
	var buf bytes.Buffer
	_, err := buf.WriteString(value)
	return buf, err
}

func convBytesToBuffer(value []byte) (bytes.Buffer, error) {
	var buf bytes.Buffer
	_, err := buf.Write(value)
	return buf, err
}
