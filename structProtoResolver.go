package structproto

import (
	"fmt"
	"reflect"

	"github.com/Bofry/structproto/tagresolver"
)

type StructProtoResolver struct {
	tagName     string
	tagResolver TagResolver

	checkDuplicateNames bool
}

func NewStructProtoResolver(option *StructProtoResolveOption) *StructProtoResolver {
	if option == nil {
		panic("specified argument 'option' cannot be nil")
	}

	r := &StructProtoResolver{
		tagName:     option.TagName,
		tagResolver: option.TagResolver,

		checkDuplicateNames: option.CheckDuplicateNames,
	}

	// use StdTagResolver if missing
	// - or -
	// use NoneTagResolver if unassign both TagResolver and TagName
	if r.tagResolver == nil {
		if len(r.tagName) > 0 {
			r.tagResolver = tagresolver.StdTagResolver
		} else {
			r.tagResolver = tagresolver.NoneTagResolver
		}
	}
	return r
}

func (r *StructProtoResolver) Resolve(target interface{}) (*Struct, error) {
	var rv reflect.Value
	switch target.(type) {
	case reflect.Value:
		rv = target.(reflect.Value)
	default:
		rv = reflect.ValueOf(target)
	}

	if !rv.IsValid() {
		return nil, fmt.Errorf("specified argument 'target' is invalid")
	}

	for i := 0; true; i++ {
		if i >= 32 {
			return nil, fmt.Errorf("exceed maximum processing calls")
		}
		switch rv.Kind() {
		case reflect.Struct:
			info, err := r.internalResolve(rv)
			if err != nil {
				return nil, err
			}
			return info, nil
		case reflect.Ptr:
			if rv.IsNil() {
				rv = reflect.New(rv.Type().Elem())
			}
			rv = rv.Elem()
		default:
			return nil, fmt.Errorf("specified argument 'target' must be pointer to struct")
		}
	}
	return nil, nil
}

func (r *StructProtoResolver) internalResolve(rv reflect.Value) (*Struct, error) {
	var prototype = makeStruct(rv)
	t := rv.Type()
	count := t.NumField()
	for i := 0; i < count; i++ {
		fieldname := t.Field(i).Name
		token := r.getTagContent(t.Field(i))
		tag, err := r.tagResolver(fieldname, token)
		if err != nil {
			return nil, err
		}
		if tag != nil {
			field := &FieldInfoImpl{
				idName: fieldname,
				name:   tag.Name,
				index:  i,
				desc:   tag.Desc,
				tag:    t.Field(i).Tag,
			}
			field.appendFlags(tag.Flags...)

			if r.checkDuplicateNames {
				_, ok := prototype.fields[tag.Name]
				if ok {
					return nil, fmt.Errorf("find duplicate name '%s' on field '%s'", tag.Name, fieldname)
				}
			}
			prototype.fields[tag.Name] = field
			if field.HasFlag(RequiredFlag) {
				prototype.requiredFields.append(tag.Name)
			}
		}
	}
	return prototype, nil
}

func (r *StructProtoResolver) getTagContent(field reflect.StructField) string {
	if len(r.tagName) > 0 {
		return field.Tag.Get(r.tagName)
	}
	return ""
}
