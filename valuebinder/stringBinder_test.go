package valuebinder

import (
	"bytes"
	"encoding/json"
	"net"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/Bofry/types"
)

func TestStringBinder_WithBool(t *testing.T) {
	var target bool = false
	var input = "true"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}
	if target != true {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", true, target)
	}
}

func TestStringBinder_WithInt(t *testing.T) {
	var target int = 0
	var input = "1"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}
	if target != 1 {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", 1, target)
	}
}

func TestStringBinder_WithString(t *testing.T) {
	var target string = ""
	var input = "foo"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	binder.Bind(input)
	if target != "foo" {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", "foo", target)
	}
}

func TestStringBinder_WithDuration(t *testing.T) {
	var target time.Duration
	var input = "1m2s"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected, _ := time.ParseDuration("1m2s")
	if target != expected {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringBinder_WithStringArray(t *testing.T) {
	var target []string
	var input = "alice,bob,carlo,david,frank,george"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := []string{"alice", "bob", "carlo", "david", "frank", "george"}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringBinder_WithIntArray(t *testing.T) {
	var target []int
	var input = "1,1,2,3,5,8,13"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := []int{1, 1, 2, 3, 5, 8, 13}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringBinder_WithRawContent(t *testing.T) {
	var target types.RawContent
	var input = "binary content"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := types.RawContent("binary content")
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringBinder_WithRawMessage(t *testing.T) {
	var target json.RawMessage
	var input = "binary content"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := json.RawMessage("binary content")
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringBinder_WithIP(t *testing.T) {
	var target net.IP
	var input = "192.168.56.12"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := net.ParseIP("192.168.56.12")
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringBinder_WithIPArray(t *testing.T) {
	var target []net.IP
	var input = "192.168.56.12,192.168.56.16"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	expected := []net.IP{net.ParseIP("192.168.56.12"), net.ParseIP("192.168.56.16")}
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringBinder_WithURL(t *testing.T) {
	var target url.URL
	var input = "http://localhost:80/path/"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	url, _ := url.Parse("http://localhost:80/path/")
	expected := *url
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}

func TestStringBinder_WithBuffer(t *testing.T) {
	var target bytes.Buffer
	var input = "The quick brown fox jumps over the lazy dog"

	rv := reflect.ValueOf(&target).Elem()
	binder := StringBinder(rv)
	err := binder.Bind(input)
	if err != nil {
		t.Error(err)
	}

	var buf bytes.Buffer
	buf.WriteString("The quick brown fox jumps over the lazy dog")
	expected := buf
	if !reflect.DeepEqual(target, expected) {
		t.Errorf("assert 'target':: expected '%#v', got '%#v'", expected, target)
	}
}
