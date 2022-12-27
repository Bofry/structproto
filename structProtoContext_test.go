package structproto

import (
	"reflect"
	"testing"
	"time"
)

func TestStructProtoContext(t *testing.T) {
	c := struct {
		Name       string    `demo:"*NAME"`
		Age        *int      `demo:"*AGE"`
		Alias      []string  `demo:"ALIAS"`
		DatOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
		Remark     string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
		Numbers    []int     `demo:"NUMBERS"`
	}{}

	prototype, err := Prototypify(&c, &StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		t.Error(err)
	}

	context := buildStructProtoContext(prototype)

	expectedFieldNames := []string{"NAME", "AGE", "ALIAS", "DATE_OF_BIRTH", "REMARK", "NUMBERS"}
	if !reflect.DeepEqual(expectedFieldNames, context.FieldNames()) {
		t.Errorf("assert 'structprotoContext.AllFields()':: expected '%#v', got '%#v'", expectedFieldNames, context.FieldNames())
	}
	expectedRequiredFields := []string{"AGE", "NAME"}
	if !reflect.DeepEqual(expectedRequiredFields, context.RequiredFields()) {
		t.Errorf("assert 'structprotoContext.AllRequiredFields()':: expected '%#v', got '%#v'", expectedRequiredFields, context.RequiredFields())
	}

	{
		field := context.getFieldInfoImpl("NAME")
		if field == nil {
			t.Errorf("assert 'FieldInfo.Field(\"NAME\")':: expected not nil, got '%#v'", field)
		}
		expectedIDName := "Name"
		if field.IDName() != expectedIDName {
			t.Errorf("assert 'FieldInfo.IDName()':: expected '%#v', got '%#v'", expectedIDName, field.IDName())
		}
		expectedName := "NAME"
		if field.Name() != expectedName {
			t.Errorf("assert 'FieldInfo.Name()':: expected '%#v', got '%#v'", expectedName, field.Name())
		}
		expectedIndex := 0
		if field.Index() != expectedIndex {
			t.Errorf("assert 'FieldInfo.Index()':: expected '%#v', got '%#v'", expectedIndex, field.Index())
		}
		expectedFlags := []string{"required"}
		if !reflect.DeepEqual(expectedFlags, field.flags.ToArray()) {
			t.Errorf("assert 'FieldInfo.Flags()':: expected '%#v', got '%#v'", expectedFlags, field.flags.ToArray())
		}
	}

	if !context.IsRequired("NAME") {
		t.Errorf("assert 'structprotoContext.IsRequiredField(\"NAME\")':: expected '%#v', got '%#v'", expectedRequiredFields, context.IsRequired("NAME"))
	}
	if context.IsRequired("unknown") {
		t.Errorf("assert 'structprotoContext.IsRequiredField(\"unknown\")':: expected '%#v', got '%#v'", expectedRequiredFields, context.IsRequired("unknown"))
	}

	// TODO: test context.ChechIfMissingRequireFields
}
