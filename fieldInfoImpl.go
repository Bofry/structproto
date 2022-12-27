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

func (f *FieldInfoImpl) IDName() string {
	return f.idName
}

func (f *FieldInfoImpl) Name() string {
	return f.name
}

func (f *FieldInfoImpl) Desc() string {
	return f.desc
}

func (f *FieldInfoImpl) Index() int {
	return f.index
}

func (f *FieldInfoImpl) FindFlag(predicate func(v string) bool) bool {
	return f.flags.find(predicate)
}

func (f *FieldInfoImpl) HasFlag(v string) bool {
	return f.flags.has(v)
}

func (f *FieldInfoImpl) Tag() reflect.StructTag {
	return f.tag
}

func (f *FieldInfoImpl) appendFlags(values ...string) {
	if len(values) > 0 {
		for _, v := range values {
			if len(v) == 0 {
				continue
			}
			f.flags.append(v)
		}
	}
}
