package reflecting

import (
	"reflect"
	"testing"

	_ "unsafe"

	_ "github.com/cstockton/go-conv"
)

//go:linkname indirectVal github.com/cstockton/go-conv/internal/refutil.IndirectVal
func indirectVal(val reflect.Value) reflect.Value

func TestAssignZero_String(t *testing.T) {
	var source string

	in := reflect.ValueOf(&source)
	{
		expectIsValid := true
		if in.IsValid() != expectIsValid {
			t.Errorf("in.IsValid() expect: %v, got: %v\n", expectIsValid, in.IsValid())
		}
		expectType := "*string"
		if in.Type().String() != expectType {
			t.Errorf("in.Type() expect: %v, got: %v\n", expectType, in.Type())
		}
	}
	out := AssignZero(in)
	{
		expectIsValid := true
		if out.IsValid() != expectIsValid {
			t.Errorf("out.IsValid() expect: %v, got: %v\n", expectIsValid, out.IsValid())
		}
		expectType := "*string"
		if out.Type().String() != expectType {
			t.Errorf("out.Type() expect: %v, got: %v\n", expectType, out.Type())
		}
	}

	scalar := indirectVal(out)
	scalar.SetString("foo")
	{
		expectSource := "foo"
		if expectSource != source {
			t.Errorf("source expect: %v, got: %v\n", expectSource, source)
		}
	}
}

func TestAssignZero_PtrString(t *testing.T) {
	var source *string

	in := reflect.ValueOf(&source)
	{
		expectIsValid := true
		if in.IsValid() != expectIsValid {
			t.Errorf("in.IsValid() expect: %v, got: %v\n", expectIsValid, in.IsValid())
		}
		expectType := "**string"
		if in.Type().String() != expectType {
			t.Errorf("in.Type() expect: %v, got: %v\n", expectType, in.Type())
		}
	}
	out := AssignZero(in)
	{
		expectIsValid := true
		if out.IsValid() != expectIsValid {
			t.Errorf("out.IsValid() expect: %v, got: %v\n", expectIsValid, out.IsValid())
		}
		expectType := "**string"
		if out.Type().String() != expectType {
			t.Errorf("out.Type() expect: %v, got: %v\n", expectType, out.Type())
		}
	}

	scalar := indirectVal(out)
	scalar.SetString("foo")
	{
		expectSource := "foo"
		if expectSource != *source {
			t.Errorf("source expect: %v, got: %v\n", expectSource, source)
		}
	}
}

func TestAssignZero_PtrPtrString(t *testing.T) {
	var source **string

	in := reflect.ValueOf(&source)
	{
		expectIsValid := true
		if in.IsValid() != expectIsValid {
			t.Errorf("in.IsValid() expect: %v, got: %v\n", expectIsValid, in.IsValid())
		}
		expectType := "***string"
		if in.Type().String() != expectType {
			t.Errorf("in.Type() expect: %v, got: %v\n", expectType, in.Type())
		}
	}
	out := AssignZero(in)
	{
		expectIsValid := true
		if out.IsValid() != expectIsValid {
			t.Errorf("out.IsValid() expect: %v, got: %v\n", expectIsValid, out.IsValid())
		}
		expectType := "***string"
		if out.Type().String() != expectType {
			t.Errorf("out.Type() expect: %v, got: %v\n", expectType, out.Type())
		}
	}

	scalar := indirectVal(out)
	scalar.SetString("foo")
	{
		expectSource := "foo"
		if expectSource != **source {
			t.Errorf("source expect: %v, got: %v\n", expectSource, source)
		}
	}
}

func TestAssignZero_PtrPtrPtrString(t *testing.T) {
	var source ***string

	in := reflect.ValueOf(&source)
	{
		expectIsValid := true
		if in.IsValid() != expectIsValid {
			t.Errorf("in.IsValid() expect: %v, got: %v\n", expectIsValid, in.IsValid())
		}
		expectType := "****string"
		if in.Type().String() != expectType {
			t.Errorf("in.Type() expect: %v, got: %v\n", expectType, in.Type())
		}
	}
	out := AssignZero(in)
	{
		expectIsValid := true
		if out.IsValid() != expectIsValid {
			t.Errorf("out.IsValid() expect: %v, got: %v\n", expectIsValid, out.IsValid())
		}
		expectType := "****string"
		if out.Type().String() != expectType {
			t.Errorf("out.Type() expect: %v, got: %v\n", expectType, out.Type())
		}
	}

	scalar := indirectVal(out)
	scalar.SetString("foo")
	{
		expectSource := "foo"
		if expectSource != ***source {
			t.Errorf("source expect: %v, got: %v\n", expectSource, source)
		}
	}
}
