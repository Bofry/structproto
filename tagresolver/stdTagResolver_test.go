package tagresolver

import (
	"reflect"
	"testing"
)

func TestStdTagResolver(t *testing.T) {
	{
		tag, err := StdTagResolver("name", "real_name,required,nonempty;the customer's real name")
		if err != nil {
			t.Errorf("should not error, but got %v", err)
		} else {
			var expectedName string = "real_name"
			if tag.Name != expectedName {
				t.Errorf("assert Tag.Name expected '%+v', got '%+v'", expectedName, tag.Name)
			}
			var expectedFlags []string = []string{"required", "nonempty"}
			if !reflect.DeepEqual(tag.Flags, expectedFlags) {
				t.Errorf("assert Tag.Flags expected '%+v', got '%+v'", expectedFlags, tag.Flags)
			}
			var expectedDesc string = "the customer's real name"
			if tag.Desc != expectedDesc {
				t.Errorf("assert Tag.Desc expected '%+v', got '%+v'", expectedDesc, tag.Desc)
			}
		}
	}

	{
		tag, err := StdTagResolver("name", "*real_name,required,nonempty;the customer's real name")
		if err != nil {
			t.Errorf("should not error, but got %v", err)
		} else {
			var expectedName string = "*real_name"
			if tag.Name != expectedName {
				t.Errorf("assert Tag.Name expected '%+v', got '%+v'", expectedName, tag.Name)
			}
			var expectedFlags []string = []string{"required", "nonempty"}
			if !reflect.DeepEqual(tag.Flags, expectedFlags) {
				t.Errorf("assert Tag.Flags expected '%+v', got '%+v'", expectedFlags, tag.Flags)
			}
			var expectedDesc string = "the customer's real name"
			if tag.Desc != expectedDesc {
				t.Errorf("assert Tag.Desc expected '%+v', got '%+v'", expectedDesc, tag.Desc)
			}
		}
	}
}
