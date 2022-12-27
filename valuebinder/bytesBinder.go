package valuebinder

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/Bofry/structproto/common"
	"github.com/Bofry/structproto/util/reflectutil"
)

var (
	_ common.ValueBindProvider = BuildBytesBinder
	_ common.ValueBinder       = new(BytesBinder)
)

type BytesBinder reflect.Value

func BuildBytesBinder(rv reflect.Value) common.ValueBinder {
	return BytesBinder(rv)
}

func (binder BytesBinder) Bind(input interface{}) error {
	buf, ok := input.([]byte)
	if !ok {
		return fmt.Errorf("cannot bind type %T from input", input)
	}
	rv := reflect.Value(binder)
	if typeOfBytes.AssignableTo(rv.Type()) {
		rv.Set(reflect.ValueOf(buf))
		return nil
	}
	return binder.bindValueImpl(rv, buf)
}

func (binder BytesBinder) bindValueImpl(rv reflect.Value, v []byte) error {
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
		return StringBinder(reflect.Value(binder)).bindValueImpl(rv, str)
	}
	return err
}
