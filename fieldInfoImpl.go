package structproto

import "reflect"

var _ FieldInfo = new(FieldInfoImpl)

type FieldInfoImpl struct {
	idName string
	name   string
	desc   string
	index  int
	flags  FieldFlagSet
	tag    reflect.StructTag
}

// IDName implements FieldInfo.
func (f *FieldInfoImpl) IDName() string {
	return f.idName
}

// Name implements FieldInfo.
func (f *FieldInfoImpl) Name() string {
	return f.name
}

// Desc implements FieldInfo.
func (f *FieldInfoImpl) Desc() string {
	return f.desc
}

// Index implements FieldInfo.
func (f *FieldInfoImpl) Index() int {
	return f.index
}

// FindFlag implements FieldInfo.
func (f *FieldInfoImpl) FindFlag(predicate func(v string) bool) bool {
	return f.flags.find(predicate)
}

// HasFlag implements FieldInfo.
func (f *FieldInfoImpl) HasFlag(v string) bool {
	return f.flags.has(v)
}

// Tag implements FieldInfo.
func (f *FieldInfoImpl) Tag() reflect.StructTag {
	return f.tag
}

func (f *FieldInfoImpl) appendFlags(values ...string) {
	if len(values) == 0 {
		return
	}

	for _, v := range values {
		if len(v) == 0 {
			continue
		}
		f.flags.append(v)
	}
}
