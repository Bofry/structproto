package valuebinder

import (
	"fmt"
	"reflect"

	"github.com/Bofry/structproto/common"
	"github.com/Bofry/structproto/valuebinder/converter"
	"github.com/cstockton/go-conv"
)

var (
	knownTypeBinderTable = map[reflect.Type]typeBinder{
		typeOfDuration:   bindDuration,
		typeOfRawContent: bindRawContent,
		typeOfRawMessage: bindRawMessage,
		typeOfIP:         bindIP,
		typeOfUrl:        bindUrl,
		typeOfTime:       bindTime,
		typeOfBuffer:     bindBuffer,
	}

	errBindingUnsupportedType = fmt.Errorf("cannot bind specified type")
)

func bindKnownType(rv reflect.Value, v interface{}) (bool, error) {
	if reflect.PointerTo(rv.Type()).Implements(typeOfUnmarshaler) {
		u := reflect.New(rv.Type())
		err := bindUnmarshaler(u, v)
		if err != nil {
			return true, err
		}
		rv.Set(reflect.Indirect(u))
		return true, nil
	}

	if binder, ok := knownTypeBinderTable[rv.Type()]; ok {
		return true, binder(rv, v)
	}
	return false, nil
}

func bindValue(rv reflect.Value, v interface{}) error {
	switch rv.Kind() {
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
		int, err := conv.Int64(v)
		if err != nil {
			return &ValueBindingError{v, rv.Kind().String(), err}
		}
		rv.SetInt(int)
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
	default:
		return &ValueBindingError{v, rv.Kind().String(), errBindingUnsupportedType}
	}
	return nil
}

func bindDuration(rv reflect.Value, v interface{}) error {
	duration, err := conv.Duration(v)
	if err != nil {
		return &ValueBindingError{v, rv.Type().String(), err}
	}
	rv.Set(reflect.ValueOf(duration))
	return nil
}

func bindRawContent(rv reflect.Value, v interface{}) error {
	content, err := converter.RawContent(v)
	if err != nil {
		return &ValueBindingError{v, rv.Type().String(), err}
	}
	rv.Set(reflect.ValueOf(content))
	return nil
}

func bindRawMessage(rv reflect.Value, v interface{}) error {
	message, err := converter.RawMessage(v)
	if err != nil {
		return &ValueBindingError{v, rv.Type().String(), err}
	}
	rv.Set(reflect.ValueOf(message))
	return nil
}

func bindIP(rv reflect.Value, v interface{}) error {
	ip, err := converter.IP(v)
	if err != nil {
		return &ValueBindingError{v, rv.Type().String(), err}
	}
	rv.Set(reflect.ValueOf(ip))
	return nil
}

func bindUrl(rv reflect.Value, v interface{}) error {
	url, err := converter.Url(v)
	if err != nil {
		return &ValueBindingError{v, rv.Type().String(), err}
	}
	rv.Set(reflect.ValueOf(url))
	return nil
}

func bindTime(rv reflect.Value, v interface{}) error {
	time, err := conv.Time(v)
	if err != nil {
		return &ValueBindingError{v, rv.Type().String(), err}
	}
	rv.Set(reflect.ValueOf(time))
	return nil
}

func bindBuffer(rv reflect.Value, v interface{}) error {
	buf, err := converter.Buffer(v)
	if err != nil {
		return &ValueBindingError{v, rv.Type().String(), err}
	}
	rv.Set(reflect.ValueOf(buf))
	return nil
}

func bindUnmarshaler(rv reflect.Value, v interface{}) error {
	rvUnmarshaler := rv.Convert(typeOfUnmarshaler)
	unmarshaler, ok := rvUnmarshaler.Interface().(common.Unmarshaler)
	if !ok {
		panic(fmt.Errorf("cannot convert %s to StructUnmarshaler", rv.Type().String()))
	}
	return unmarshaler.UnmarshalStruct(v)
}
