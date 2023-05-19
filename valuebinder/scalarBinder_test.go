package valuebinder

import (
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/Bofry/types"
)

func TestScalarBinder_WithBool(t *testing.T) {
	var target bool
	var input = []byte("true")

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)

	expected := true
	if target != expected {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithInt(t *testing.T) {
	var target int
	var input = []byte("1")

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)

	expected := 1
	if target != expected {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithString(t *testing.T) {
	var target string = ""
	var input = []byte("foo")

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)

	expected := "foo"
	if target != expected {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithDuration(t *testing.T) {
	var target time.Duration
	var input = "1m2s"

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected, _ := time.ParseDuration("1m2s")
	if target != expected {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithRawContent(t *testing.T) {
	var target types.RawContent
	var input = []byte("foo")

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := types.RawContent([]byte("foo"))
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithIntSlice(t *testing.T) {
	var target []int
	var input = []int{1, 2, 3}

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithStruct(t *testing.T) {
	type Model struct {
		ID    string
		Value string
	}
	var target Model
	var input = Model{
		ID:    "1",
		Value: "foo",
	}

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := Model{
		ID:    "1",
		Value: "foo",
	}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithStructPtr(t *testing.T) {
	type Model struct {
		ID    string
		Value string
	}
	var target *Model
	var input = &Model{
		ID:    "1",
		Value: "foo",
	}

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := &Model{
		ID:    "1",
		Value: "foo",
	}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithFunc(t *testing.T) {
	type Func func() string

	var target Func
	var input = func() string {
		return "foo"
	}

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := func() string {
		return "foo"
	}
	if target() != expected() {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected(), target())
	}
}

func TestScalarBinder_WithIP(t *testing.T) {
	var target net.IP
	var input = net.ParseIP("127.0.0.3")

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)

	expected := net.ParseIP("127.0.0.3")
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithIPviaString(t *testing.T) {
	var target net.IP
	var input = "127.0.0.3"

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)

	expected := net.ParseIP("127.0.0.3")
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithIPSlice(t *testing.T) {
	var target []net.IP
	var input = []string{"127.0.0.3", "127.0.0.4"}

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)

	expected := []net.IP{net.ParseIP("127.0.0.3"), net.ParseIP("127.0.0.4")}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithMap_SameType(t *testing.T) {
	var target map[string]int
	var input = map[string]int{
		"one": 1,
		"two": 2,
	}

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)

	expected := input
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithMap_ValueIntToString(t *testing.T) {
	var target map[string]string
	var input = map[string]int{
		"one": 1,
		"two": 2,
	}

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)

	expected := map[string]string{
		"one": "1",
		"two": "2",
	}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithMap_KeyStringToCustomStringType(t *testing.T) {
	type strkey string

	var target map[strkey]int
	var input = map[string]int{
		"one": 1,
		"two": 2,
	}

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)

	expected := map[strkey]int{
		"one": 1,
		"two": 2,
	}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestScalarBinder_WithMap_ValueIntToPtrInt(t *testing.T) {
	createIntPtr := func(v int) *int {
		return &v
	}

	var target map[string]*int
	var input = map[string]int{
		"one": 1,
		"two": 2,
	}

	rv := reflect.ValueOf(&target).Elem()
	binder := ScalarBinder(rv)
	binder.Bind(input)

	expected := map[string]*int{
		"one": createIntPtr(1),
		"two": createIntPtr(2),
	}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}
