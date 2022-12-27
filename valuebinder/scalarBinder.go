package valuebinder

import (
	"reflect"

	"github.com/Bofry/structproto/common"
	"github.com/Bofry/structproto/util/reflectutil"
	"github.com/Bofry/structproto/valuebinder/converter"
	"github.com/cstockton/go-conv"
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
	rv = reflect.Indirect(reflectutil.AssignZero(rv))
	var err error

	switch rv.Kind() {
	case reflect.Struct:
		switch rv.Type() {
		case typeOfUrl:
			url, err := converter.Url(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(url))
		case typeOfTime:
			time, err := conv.Time(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(time))
		case typeOfBuffer:
			buf, err := converter.Buffer(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(buf))
		default:
			return &ValueBindingError{v, rv.Type().String(), err}
		}
	case reflect.Bool:
		bool, err := conv.Bool(v)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetBool(bool)
	case reflect.String:
		string, err := conv.String(v)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetString(string)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch rv.Type() {
		case typeOfDuration:
			duration, err := conv.Duration(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(duration))
		default:
			int, err := conv.Int64(v)
			if err != nil {
				return &ValueBindingError{v, rv.Kind().String(), err}
			}
			rv.SetInt(int)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uint, err := conv.Uint64(v)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetUint(uint)
	case reflect.Float32, reflect.Float64:
		float, err := conv.Float64(v)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetFloat(float)
	case reflect.Array, reflect.Slice:
		switch rv.Type() {
		case typeOfRawContent:
			content, err := converter.RawContent(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(content))
		case typeOfRawMessage:
			message, err := converter.RawMessage(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(message))
		case typeOfIP:
			ip, err := converter.IP(v)
			if err != nil {
				return &ValueBindingError{v, rv.Type().String(), err}
			}
			rv.Set(reflect.ValueOf(ip))
		default:
			{
				ri := reflect.ValueOf(v)
				switch ri.Kind() {
				case reflect.Array, reflect.Slice:
					size := ri.Len()
					container := reflect.MakeSlice(rv.Type(), size, size)
					for i := 0; i < size; i++ {
						err := binder.bindValueImpl(container.Index(i), ri.Index(i).Interface())
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
			}
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
	default:
		return &ValueBindingError{v, rv.Kind().String(), err}
	}
	return err
}
