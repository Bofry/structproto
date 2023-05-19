package converter

import (
	"net"
	"reflect"
)

var (
	typeOfIP = reflect.TypeOf(net.IP(nil))
)

func IP(from interface{}) (net.IP, error) {
	if T, ok := from.(net.IP); ok {
		return T, nil
	} else if T, ok := from.(*net.IP); ok {
		return *T, nil
	} else if T, ok := from.([]byte); ok {
		return convBytesToIP(T)
	} else if T, ok := from.(string); ok {
		return convStringToIP(T)
	}

	rv := reflect.ValueOf(indirect(from))
	switch rv.Kind() {
	case reflect.String:
		return convStringToIP(rv.String())
	case reflect.Array, reflect.Slice:
		if rv.CanInterface() {
			if T, ok := rv.Interface().([]byte); ok {
				return convBytesToIP(T)
			}
		}
	}
	return nil, newConvErr(from, "net.IP")
}

func convStringToIP(value string) (net.IP, error) {
	ip := net.ParseIP(value)
	return ip, nil
}

func convBytesToIP(value []byte) (net.IP, error) {
	return value, nil
}
