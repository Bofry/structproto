package common

import "reflect"

const (
	RequiredFlag = "required"
	BlankFlag    = "_"
)

type (
	Tag struct {
		Name  string
		Flags []string
		Desc  string
	}

	Unmarshaler interface {
		UnmarshalStruct(v interface{}) error
	}

	ValueBinder interface {
		Bind(v interface{}) error
	}

	TagResolver       func(fieldname, token string) (*Tag, error)
	ValueBindProvider func(rv reflect.Value) ValueBinder
)
