package valuebinder

import (
	"reflect"

	"github.com/Bofry/structproto/common"
	"github.com/Bofry/structproto/reflecting"
)

var (
	_ common.ValueBindProvider = BuildScalarBinder
	_ common.ValueBinder       = new(ScalarBinder)
)

type ScalarBinder reflect.Value

func BuildScalarBinder(rv reflect.Value) common.ValueBinder {
	return ScalarBinder(rv)
}

func (binder ScalarBinder) Bind(v interface{}) error {
	rf := reflect.Value(binder)
	{
		rv := reflect.ValueOf(v)
		if rv.Type().AssignableTo(rf.Type()) {
			rf.Set(rv)
			return nil
		}
	}
	return binder.bindValueImpl(rf, v)
}

func (binder ScalarBinder) bindValueImpl(rv reflect.Value, v interface{}) error {
	rv = indirectVal(reflecting.AssignZero(rv))
	var err error

	if ok, err := bindKnownType(rv, v); ok {
		return err
	}

	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		in := reflect.ValueOf(v)
		switch in.Kind() {
		case reflect.Array, reflect.Slice:
			size := in.Len()
			container := reflect.MakeSlice(rv.Type(), size, size)
			for i := 0; i < size; i++ {
				err := binder.bindValueImpl(container.Index(i), in.Index(i).Interface())
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
		default:
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
	case reflect.Map:
		in := reflect.ValueOf(v)
		switch in.Kind() {
		case reflect.Map:
			out := reflect.MakeMap(rv.Type())
			iter := in.MapRange()
			for iter.Next() {
				key := iter.Key()
				val := iter.Value()

				if in.Type() != rv.Type() {
					outKey := reflect.New(out.Type().Key())
					outVal := reflect.New(out.Type().Elem())

					err = binder.bindValueImpl(outKey, key.Interface())
					if err != nil {
						return err
					}
					err = binder.bindValueImpl(outVal, val.Interface())
					if err != nil {
						return err
					}
					key = outKey.Elem()
					val = outVal.Elem()
				}
				out.SetMapIndex(key, val)
			}
			rv.Set(out)
		}
	}
	if rv.IsZero() {
		return bindValue(rv, v)
	}
	return err
}
