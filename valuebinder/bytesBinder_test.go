package valuebinder

import (
	"reflect"
	"testing"
	"time"
)

func TestBytesBinder_WithBytes(t *testing.T) {
	var v []byte
	var input = []byte("true")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesBinder(rv)
	binder.Bind(input)
	if !reflect.DeepEqual(input, v) {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", true, v)
	}
}

func TestBytesBinder_WithBool(t *testing.T) {
	var v bool = false
	var input = []byte("true")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesBinder(rv)
	binder.Bind(input)
	if v != true {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", true, v)
	}
}

func TestBytesBinder_WithInt(t *testing.T) {
	var v int = 0
	var input = []byte("1")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesBinder(rv)
	binder.Bind(input)
	if v != 1 {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", 1, v)
	}
}

func TestBytesBinder_WithString(t *testing.T) {
	var v string = ""
	var input = []byte("foo")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesBinder(rv)
	binder.Bind(input)
	if v != "foo" {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", "foo", v)
	}
}

func TestBytesBinder_WithDuration(t *testing.T) {
	var v time.Duration
	var input = []byte("1m2s")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected, _ := time.ParseDuration("1m2s")
	if v != expected {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", expected, v)
	}
}

func TestBytesBinder_WithStringArray(t *testing.T) {
	var v []string
	var input = []byte("alice,bob,carlo,david,frank,george")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := []string{"alice", "bob", "carlo", "david", "frank", "george"}
	if !reflect.DeepEqual(v, expected) {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", expected, v)
	}
}

func TestBytesBinder_WithIntArray(t *testing.T) {
	var v []int
	var input = []byte("1,1,2,3,5,8,13")

	rv := reflect.ValueOf(&v).Elem()
	binder := BytesBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := []int{1, 1, 2, 3, 5, 8, 13}
	if !reflect.DeepEqual(v, expected) {
		t.Errorf("assert 'v':: expected '%#v', got '%#v'", expected, v)
	}
}
