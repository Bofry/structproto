package valuebinder

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/Bofry/structproto/internal"
	"github.com/Bofry/structproto/util/reflectutil"
)

var (
	_ internal.ValueBindProvider = BuildBytesArgsBinder
	_ internal.ValueBinder       = new(BytesArgsBinder)
)

type BytesArgsBinder reflect.Value

func BuildBytesArgsBinder(rv reflect.Value) internal.ValueBinder {
	return BytesArgsBinder(rv)
}

func (binder BytesArgsBinder) Bind(input interface{}) error {
	buf, ok := input.([]byte)
	if !ok {
		return fmt.Errorf("cannot bind type %T from input", input)
	}
	rv := reflect.Value(binder)
	return binder.bindValueImpl(rv, buf)
}

func (binder BytesArgsBinder) bindValueImpl(rv reflect.Value, v []byte) error {
	rv = reflect.Indirect(reflectutil.AssignZero(rv))
	var err error

	kind := rv.Kind()
	typ := rv.Type()
	if kind == reflect.Struct && typ == typeOfBuffer {
		var buf bytes.Buffer
		_, err := buf.Write(v)
		if err != nil {
			return &ValueBindingError{v, rv.Type().String(), err}
		}
		rv.Set(reflect.ValueOf(buf))
	} else {
		str := string(v)
		return StringArgsBinder(reflect.Value(binder)).bindValueImpl(rv, str)
	}
	return err
}
