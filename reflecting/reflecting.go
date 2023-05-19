package reflecting

import (
	"reflect"
)

func AssignZero(rv reflect.Value) reflect.Value {
	var last = rv
	for {
		if last.Kind() != reflect.Ptr {
			break
		}

		if last.IsNil() {
			last.Set(reflect.New(last.Type().Elem()))
		}

		next := last.Elem()
		if next.Kind() != reflect.Ptr {
			break
		}
		last = next
	}
	return rv
}
