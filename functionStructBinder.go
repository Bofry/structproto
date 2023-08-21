package structproto

import (
	"reflect"
)

var _ StructBinder = new(FunctionStructBinder)

type FunctionStructBinder struct {
	mapper StructMapper
}

// Bind implements StructBinder.
func (binder *FunctionStructBinder) Bind(field FieldInfo, rv reflect.Value) error {
	return binder.mapper(field, rv)
}

// Deinit implements StructBinder.
func (binder *FunctionStructBinder) Deinit(context *StructProtoContext) error {
	return nil
}

// Init implements StructBinder.
func (binder *FunctionStructBinder) Init(context *StructProtoContext) error {
	return nil
}
