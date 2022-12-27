package structproto

import (
	"reflect"
	"testing"
	"time"
)

func TestStructInternal(t *testing.T) {
	c := struct {
		Name        string    `demo:"*NAME"`
		Age         *int      `demo:"*AGE"`
		Alias       []string  `demo:"ALIAS"`
		DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
		Remark      string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
		Numbers     []int     `demo:"NUMBERS"`
	}{}

	prototype, err := Prototypify(&c, &StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		t.Error(err)
	}
	expectedFields := map[string]FieldInfoImpl{
		"NAME": {
			idName: "Name",
			name:   "NAME",
			desc:   "",
			index:  0,
			flags:  []string{"required"},
			tag:    `demo:"*NAME"`,
		},
		"AGE": {
			idName: "Age",
			name:   "AGE",
			desc:   "",
			index:  1,
			flags:  []string{"required"},
			tag:    `demo:"*AGE"`,
		},
		"ALIAS": {
			idName: "Alias",
			name:   "ALIAS",
			desc:   "",
			index:  2,
			flags:  []string(nil),
			tag:    `demo:"ALIAS"`,
		},
		"DATE_OF_BIRTH": {
			idName: "DateOfBirth",
			name:   "DATE_OF_BIRTH",
			desc:   "the character's birth of date",
			index:  3,
			flags:  []string(nil),
			tag:    `demo:"DATE_OF_BIRTH;the character's birth of date"`,
		},
		"REMARK": {
			idName: "Remark",
			name:   "REMARK",
			desc:   "note the character's personal favor",
			index:  4,
			flags:  []string{"flag1", "flag2"},
			tag:    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`,
		},
		"NUMBERS": {
			idName: "Numbers",
			name:   "NUMBERS",
			desc:   "",
			index:  5,
			flags:  []string(nil),
			tag:    `demo:"NUMBERS"`,
		},
	}

	if len(expectedFields) != len(prototype.fields) {
		t.Errorf("assert 'structproto.fields' length :: expected '%v', got '%v'", len(expectedFields), len(prototype.fields))
	}
	for k, v := range expectedFields {
		if f, ok := prototype.fields[k]; !ok {
			t.Errorf("assert 'structproto.fields' key '%s' not found", k)
		} else {
			if (f.IDName() != v.idName) ||
				(f.Name() != v.name) ||
				(f.Index() != v.index) ||
				(f.Desc() != v.desc) ||
				(f.Tag() != v.tag) ||
				(!reflect.DeepEqual([]string(f.flags), []string(v.flags))) {
				t.Errorf("assert 'structproto.fields' key '%s' :: expected '%#v', got '%#v'", k, v, f)
			}
		}
	}
	expectedRequiredFields := FieldFlagSet([]string{"AGE", "NAME"})
	if !reflect.DeepEqual(expectedRequiredFields, prototype.requiredFields) {
		t.Errorf("assert 'mockCharacter.requiredFields':: expected '%#v', got '%#v'", expectedRequiredFields, prototype.requiredFields)
	}
}
