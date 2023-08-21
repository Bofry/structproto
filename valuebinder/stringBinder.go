package valuebinder

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Bofry/structproto/common"
	"github.com/Bofry/structproto/reflecting"
)

var (
	_ common.ValueBindProvider = BuildStringBinder
	_ common.ValueBinder       = new(StringBinder)
)

type StringBinder reflect.Value

func BuildStringBinder(rv reflect.Value) common.ValueBinder {
	return StringBinder(rv)
}

func (binder StringBinder) Bind(input interface{}) error {
	v, ok := input.(string)
	if !ok {
		return fmt.Errorf("cannot bind type %T from input", input)
	}
	rv := reflect.Value(binder)

	if typeOfString.AssignableTo(rv.Type()) {
		rv.Set(reflect.ValueOf(v))
		return nil
	}
	return binder.bindValueImpl(rv, v)
}

func (binder StringBinder) bindValueImpl(rv reflect.Value, v string) error {
	rv = indirectVal(reflecting.AssignZero(rv))
	var err error

	if ok, err := bindKnownType(rv, v); ok {
		return err
	}

	switch rv.Kind() {
	case reflect.String:
		rv.SetString(v)
	case reflect.Array, reflect.Slice:
		if len(v) > 0 {
			array := strings.Split(v, ",")
			size := len(array)
			container := reflect.MakeSlice(rv.Type(), size, size)
			for i, elem := range array {
				err := binder.bindValueImpl(container.Index(i), elem)
				if err != nil {
					return &SliceBindingError{
						Value: v,
						Kind:  rv.Kind().String(),
						Index: i,
						Err:   err,
					}
				}
			}
			rv.Set(container)
		}
	default:
		err = bindValue(rv, v)
	}
	return err
}
