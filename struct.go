package structproto

import (
	"fmt"
	"reflect"
)

type Struct struct {
	target reflect.Value

	fields         map[string]*FieldInfoImpl
	requiredFields FieldFlagSet
}

func (s *Struct) Bind(binder StructBinder) error {
	if binder == nil {
		panic("specified argument 'binder' cannot be nil")
	}

	var (
		context = buildStructProtoContext(s)

		err error
	)

	if err = binder.Init(context); err != nil {
		return err
	}

	// bind all fields
	for _, field := range s.fields {
		err := binder.Bind(field, s.target.Field(field.index))
		if err != nil {
			return err
		}
	}

	if err = binder.Deinit(context); err != nil {
		return err
	}

	return nil
}

func (s *Struct) BindMap(values map[string]interface{}, buildValueBinder ValueBindProvider) error {
	if s == nil {
		return nil
	}

	return s.BindIterator(FieldValueMap(values), buildValueBinder)
}

func (s *Struct) BindIterator(iterator Iterator, buildValueBinder ValueBindProvider) error {
	if s == nil {
		return nil
	}
	if buildValueBinder == nil {
		return fmt.Errorf("missing ValueBinderProvider")
	}

	return s.BindChan(iterator.Iterate(), buildValueBinder)
}

func (s *Struct) BindFields(values []FieldValueEntity, buildValueBinder ValueBindProvider) error {
	if s == nil {
		return nil
	}
	if buildValueBinder == nil {
		return fmt.Errorf("missing ValueBinderProvider")
	}
	var requiredFields = s.requiredFields.clone()

	// mapping values
	for _, v := range values {
		field, val := v.Field, v.Value
		if val != nil {
			binder := s.makeFieldBinder(s.target, field, buildValueBinder)
			if binder != nil {
				err := binder.Bind(val)
				if err != nil {
					return &FieldBindingError{field, val, err}
				}

				index := requiredFields.indexOf(field)
				if index != -1 {
					// eliminate the field from slice if found
					requiredFields.removeIndex(index)
				}
			}
		}
	}

	// check if the requiredFields still have fields don't be set
	if !requiredFields.isEmpty() {
		field, _ := requiredFields.get(0)
		return &MissingRequiredFieldError{field, nil}
	}

	return nil
}

func (s *Struct) BindChan(iterator <-chan FieldValueEntity, buildValueBinder ValueBindProvider) error {
	if s == nil {
		return nil
	}
	if buildValueBinder == nil {
		return fmt.Errorf("missing ValueBinderProvider")
	}
	var requiredFields = s.requiredFields.clone()

	// mapping values
	for v := range iterator {
		field, val := v.Field, v.Value
		if val != nil {
			binder := s.makeFieldBinder(s.target, field, buildValueBinder)
			if binder != nil {
				err := binder.Bind(val)
				if err != nil {
					return &FieldBindingError{field, val, err}
				}

				index := requiredFields.indexOf(field)
				if index != -1 {
					// eliminate the field from slice if found
					requiredFields.removeIndex(index)
				}
			}
		}
	}

	// check if the requiredFields still have fields don't be set
	if !requiredFields.isEmpty() {
		field, _ := requiredFields.get(0)
		return &MissingRequiredFieldError{field, nil}
	}

	return nil
}

func (s *Struct) Map(mapper StructMapper) error {
	binder := &FunctionStructBinder{
		mapper: mapper,
	}
	return s.Bind(binder)
}

func (s *Struct) Visit(visitor StructVisitor) {
	for name, info := range s.fields {
		elem := s.target.Field(info.index)
		visitor(name, elem, info)
	}
}

func (s *Struct) makeFieldBinder(rv reflect.Value, name string, buildValueBinder ValueBindProvider) ValueBinder {
	if f, ok := s.fields[name]; ok {
		binder := buildValueBinder(rv.Field(f.index))
		return binder
	}
	return nil
}

func makeStruct(value reflect.Value) *Struct {
	var expectedFields int
	if value.Kind() == reflect.Struct {
		expectedFields = value.Type().NumField()
	}

	prototype := Struct{
		target: value,
		fields: make(map[string]*FieldInfoImpl, expectedFields),
	}
	return &prototype
}
