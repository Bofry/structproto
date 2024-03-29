package structproto

import "reflect"

type StructProtoContext Struct

func (ctx *StructProtoContext) Target() reflect.Value {
	return ctx.target
}

func (ctx *StructProtoContext) FieldInfo(name string) FieldInfo {
	return ctx.getFieldInfoImpl(name)
}

func (ctx *StructProtoContext) Field(name string) (v reflect.Value, ok bool) {
	info := ctx.FieldInfo(name)
	if info != nil {
		return ctx.target.Field(info.Index()), true
	}
	return reflect.Value{}, false
}

func (ctx *StructProtoContext) FieldNames() []string {
	var fields []string = make([]string, len(ctx.fields))
	for _, v := range ctx.fields {
		fields[v.index] = v.name
	}
	return fields
}

func (ctx *StructProtoContext) RequiredFields() []string {
	return ctx.requiredFields
}

func (ctx *StructProtoContext) IsRequired(name string) bool {
	field := ctx.getFieldInfoImpl(name)
	if field != nil {
		return field.HasFlag(RequiredFlag)
	}
	return false
}

func (ctx *StructProtoContext) CheckIfMissingRequiredFields(visitFieldProc func() <-chan string) error {
	if ctx.requiredFields.isEmpty() {
		return nil
	}

	var requiredFields = ctx.requiredFields.clone()

	for field := range visitFieldProc() {
		index := requiredFields.indexOf(field)
		if index != -1 {
			requiredFields.removeIndex(index)
		}

		// break loop if no more required fields
		if requiredFields.isEmpty() {
			return nil
		}
	}

	if !requiredFields.isEmpty() {
		field, _ := requiredFields.get(0)
		return &MissingRequiredFieldError{field, nil}
	}
	return nil
}

func (ctx *StructProtoContext) getFieldInfoImpl(name string) *FieldInfoImpl {
	if field, ok := ctx.fields[name]; ok {
		return field
	}
	return nil
}

func buildStructProtoContext(s *Struct) *StructProtoContext {
	context := StructProtoContext(*s)
	return &context
}
