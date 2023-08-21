package structproto

import (
	"reflect"

	"github.com/Bofry/structproto/common"
)

const (
	RequiredFlag = common.RequiredFlag
	BlankFlag    = common.BlankFlag
)

type (
	ValueBindProvider = common.ValueBindProvider
	ValueBinder       = common.ValueBinder
	TagResolver       = common.TagResolver
	Tag               = common.Tag

	FieldValueEntity struct {
		Field string
		Value interface{}
	}

	Iterator interface {
		Iterate() <-chan FieldValueEntity
	}

	FieldInfo interface {
		IDName() string
		Name() string
		Desc() string
		Index() int
		FindFlag(predicate func(v string) bool) bool
		HasFlag(v string) bool
		Tag() reflect.StructTag
	}

	StructBinder interface {
		Init(context *StructProtoContext) error
		Bind(field FieldInfo, rv reflect.Value) error
		Deinit(context *StructProtoContext) error
	}

	StructProtoResolveOption struct {
		TagName             string
		TagResolver         TagResolver
		CheckDuplicateNames bool
	}

	StructVisitor func(name string, rv reflect.Value, info FieldInfo)
	StructMapper  func(field FieldInfo, rv reflect.Value) error
)
