package structproto_test

import (
	"reflect"

	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/valuebinder"
)

var _ structproto.StructBinder = new(MapBinder)

type MapBinder struct {
	values map[string]string
}

func (b *MapBinder) Init(context *structproto.StructProtoContext) error {
	return nil
}

func (b *MapBinder) Bind(field structproto.FieldInfo, rv reflect.Value) error {
	name := field.Name()
	if v, ok := b.values[name]; ok {
		return valuebinder.StringBinder(rv).Bind(v)
	}
	return nil
}

func (b *MapBinder) Deinit(context *structproto.StructProtoContext) error {
	return context.CheckIfMissingRequiredFields(func() <-chan string {
		c := make(chan string)
		go func() {
			for k := range b.values {
				c <- k
			}
			close(c)
		}()
		return c
	})
}

// -------------------------------------

func ResolvePathTag(fieldname, token string) (*structproto.Tag, error) {
	var tag *structproto.Tag
	if len(token) > 0 {
		if token != "-" {
			tag = &structproto.Tag{
				Name: token,
			}
		}
	}
	return tag, nil
}

// -------------------------------------

var _ structproto.Iterator = EntitySet(nil)

type EntitySet [][2]string

func (set EntitySet) Iterate() <-chan structproto.FieldValueEntity {
	c := make(chan structproto.FieldValueEntity, 1)
	go func() {
		for _, v := range set {
			c <- structproto.FieldValueEntity{
				Field: v[0],
				Value: v[1],
			}
		}
		close(c)
	}()
	return c
}
