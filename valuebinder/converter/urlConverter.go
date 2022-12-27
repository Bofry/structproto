package converter

import (
	"net/url"
	"reflect"

	reflectutil "github.com/Bofry/structproto/util/reflectutil"
)

var (
	typeOfUrl = reflect.TypeOf(url.URL{})
)

func Url(from interface{}) (url.URL, error) {
	if T, ok := from.(url.URL); ok {
		return T, nil
	} else if T, ok := from.(*url.URL); ok {
		return *T, nil
	} else if T, ok := from.(string); ok {
		return convStringToUrl(T)
	}

	rv := reflect.ValueOf(reflectutil.Indirect(from))
	switch rv.Kind() {
	case reflect.String:
		return convStringToUrl(rv.String())
	case reflect.Struct:
		if rv.Type().ConvertibleTo(typeOfUrl) {
			valueConv := rv.Convert(typeOfUrl)
			if valueConv.CanInterface() {
				return valueConv.Interface().(url.URL), nil
			}
		}
	}
	return url.URL{}, newConvErr(from, "url.URL")
}

func convStringToUrl(value string) (url.URL, error) {
	T, err := url.Parse(value)
	return *T, err
}
