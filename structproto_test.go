package structproto_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/valuebinder"
	"go.openly.dev/pointy"
)

func TestStruct_BindMap_MissingRequiredField(t *testing.T) {
	s := struct {
		Name        string    `demo:"*NAME"`
		Age         *int      `demo:"*AGE"`
		Alias       []string  `demo:"ALIAS"`
		DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
		Remark      string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
		Numbers     []int     `demo:"NUMBERS"`
	}{}

	prototype, err := structproto.Prototypify(&s, &structproto.StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		t.Error(err)
	}
	err = prototype.BindMap(map[string]interface{}{
		"NAME":          "luffy",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}, valuebinder.BuildStringBinder)
	if err == nil {
		t.Errorf("the 'Mapper.Map()' should throw '%s' error", "with missing symbol 'AGE'")
	}
}

func TestStruct_BindIterator_WithFieldValueMap(t *testing.T) {
	type (
		model struct {
			Name        string    `demo:"*NAME"`
			Age         *int      `demo:"*AGE"`
			Alias       []string  `demo:"ALIAS"`
			DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
			Remark      string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
			Numbers     []int     `demo:"NUMBERS"`
		}
	)

	s := model{}

	prototype, err := structproto.Prototypify(&s, &structproto.StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		t.Error(err)
	}
	err = prototype.BindIterator(structproto.FieldValueMap{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
		"NUMBERS":       "5,12",
	}, valuebinder.BuildStringBinder)
	if err != nil {
		t.Error(err)
	}
	expected := model{
		Name:        "luffy",
		Age:         pointy.Int(19),
		Alias:       []string{"lucy"},
		DateOfBirth: time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC),
		Numbers:     []int{5, 12},
	}
	if !reflect.DeepEqual(expected, s) {
		t.Errorf("assert:: expected '%+v', got '%+v'", expected, s)
	}
}

func TestStruct_BindIterator_Well(t *testing.T) {
	type (
		model struct {
			Name        string    `demo:"*NAME"`
			Age         *int      `demo:"*AGE"`
			Alias       []string  `demo:"ALIAS"`
			DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
			Remark      string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
			Numbers     []int     `demo:"NUMBERS"`
		}
	)

	s := model{}

	prototype, err := structproto.Prototypify(&s, &structproto.StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		t.Error(err)
	}
	err = prototype.BindIterator(EntitySet{
		{"NAME", "luffy"},
		{"AGE", "19"},
		{"ALIAS", "lucy"},
		{"DATE_OF_BIRTH", "2020-05-05T00:00:00Z"},
		{"NUMBERS", "5,12"},
	}, valuebinder.BuildStringBinder)
	if err != nil {
		t.Error(err)
	}
	expected := model{
		Name:        "luffy",
		Age:         pointy.Int(19),
		Alias:       []string{"lucy"},
		DateOfBirth: time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC),
		Numbers:     []int{5, 12},
	}
	if !reflect.DeepEqual(expected, s) {
		t.Errorf("assert:: expected '%+v', got '%+v'", expected, s)
	}
}

func TestStruct_Bind_MissingRequiredField(t *testing.T) {
	input := map[string]string{
		"NAME": "luffy",
		// "AGE":           "19",    -- we won't set the required field
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}
	binder := &MapBinder{
		values: input,
	}

	s := struct {
		Name        string    `demo:"*NAME"`
		Age         *int      `demo:"*AGE"`
		Alias       []string  `demo:"ALIAS"`
		DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
		Remark      string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
		Numbers     []int     `demo:"NUMBERS"`
	}{}

	prototype, err := structproto.Prototypify(&s, &structproto.StructProtoResolveOption{
		TagName: "demo",
	})

	err = prototype.Bind(binder)
	if err == nil {
		t.Errorf("the 'Process()' should throw '%s' error", "missing required symbol 'AGE'")
	} else {
		missingRequiredFieldError, ok := err.(*structproto.MissingRequiredFieldError)
		if !ok {
			t.Errorf("the error expected '%T', got '%T'", &structproto.MissingRequiredFieldError{}, err)
		}
		if missingRequiredFieldError.Field != "AGE" {
			t.Errorf("assert 'MissingRequiredFieldError.Field':: expected '%v', got '%v'", "AGE", missingRequiredFieldError.Field)
		}
	}
}

func TestStruct_Bind_Well(t *testing.T) {
	type (
		model struct {
			Name        string    `demo:"*NAME"`
			Age         *int      `demo:"*AGE"`
			Alias       []string  `demo:"ALIAS"`
			DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
			Remark      string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
			Numbers     []int     `demo:"NUMBERS"`
		}
	)

	input := map[string]string{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}
	binder := &MapBinder{
		values: input,
	}

	s := model{}

	prototype, err := structproto.Prototypify(&s, &structproto.StructProtoResolveOption{
		TagName: "demo",
	})
	err = prototype.Bind(binder)
	if err != nil {
		t.Error(err)
	}

	expected := model{
		Name:        "luffy",
		Age:         pointy.Int(19),
		Alias:       []string{"lucy"},
		DateOfBirth: time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC),
	}
	if !reflect.DeepEqual(expected, s) {
		t.Errorf("assert:: expected '%+v', got '%+v'", expected, s)
	}
}

func TestStruct_BindMap_Well(t *testing.T) {
	type (
		model struct {
			Name        string    `demo:"*NAME"`
			Age         *int      `demo:"*AGE"`
			Alias       []string  `demo:"ALIAS"`
			DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
			Remark      string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
			Numbers     []int     `demo:"NUMBERS"`
		}
	)

	s := model{}

	prototype, err := structproto.Prototypify(&s, &structproto.StructProtoResolveOption{
		TagName: "demo",
	})
	err = prototype.BindMap(map[string]interface{}{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}, valuebinder.BuildStringBinder)
	if err != nil {
		t.Error(err)
	}

	expected := model{
		Name:        "luffy",
		Age:         pointy.Int(19),
		Alias:       []string{"lucy"},
		DateOfBirth: time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC),
	}
	if !reflect.DeepEqual(expected, s) {
		t.Errorf("assert:: expected '%+v', got '%+v'", expected, s)
	}
}

func TestStruct_Visitor(t *testing.T) {
	s := struct {
		Name        string    `demo:"*NAME"`
		Age         *int      `demo:"*AGE"`
		Alias       []string  `demo:"ALIAS"`
		DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
		Remark      string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
	}{}

	prototype, err := structproto.Prototypify(&s, &structproto.StructProtoResolveOption{
		TagName: "demo",
	})
	err = prototype.BindMap(map[string]interface{}{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}, valuebinder.BuildStringBinder)
	if err != nil {
		t.Error(err)
	}

	visitData := make(map[string]interface{})
	prototype.Visit(func(name string, rv reflect.Value, info structproto.FieldInfo) {
		visitData[name] = rv.Interface()
	})

	var (
		expectedAge int = 19
	)
	expectedVisitData := map[string]interface{}{
		"NAME":          "luffy",
		"AGE":           &expectedAge,
		"ALIAS":         []string{"lucy"},
		"DATE_OF_BIRTH": time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC),
		"REMARK":        "",
	}
	if len(expectedVisitData) != len(visitData) {
		t.Errorf("assert visitData size:: expected %d, got %d", len(expectedVisitData), len(visitData))
	}
	// check assigned fields
	for k, v := range visitData {
		switch k {
		case "NAME":
			expected := expectedVisitData[k].(string)
			actual, ok := v.(string)
			if !ok {
				t.Errorf("assert '%s':: expected (%T), got (%T)", k, expectedVisitData[k], v)
			}
			if actual != expected {
				t.Errorf("assert '%s':: expected '%#v', got '%#v'", k, expected, actual)
			}
		case "AGE":
			expected := expectedVisitData[k].(*int)
			actual, ok := v.(*int)
			if !ok {
				t.Errorf("assert '%s':: expected (%T), got (%T)", k, expectedVisitData[k], v)
			}
			if *actual != *expected {
				t.Errorf("assert '%s':: expected '%#v', got '%#v'", k, expected, actual)
			}
		case "ALIAS":
			expected := expectedVisitData[k].([]string)
			actual, ok := v.([]string)
			if !ok {
				t.Errorf("assert '%s':: expected (%T), got (%T)", k, expectedVisitData[k], v)
			}
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("assert '%s':: expected '%#v', got '%#v'", k, expected, actual)
			}
		case "DATE_OF_BIRTH":
			expected := expectedVisitData[k].(time.Time)
			actual, ok := v.(time.Time)
			if !ok {
				t.Errorf("assert '%s':: expected (%T), got (%T)", k, expectedVisitData[k], v)
			}
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("assert '%s':: expected '%#v', got '%#v'", k, expected, actual)
			}
		case "REMARK":
			expected := expectedVisitData[k].(string)
			actual, ok := v.(string)
			if !ok {
				t.Errorf("assert '%s':: expected (%T), got (%T)", k, expectedVisitData[k], v)
			}
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("assert '%s':: expected '%#v', got '%#v'", k, expected, actual)
			}
		default:
			t.Errorf("unknown field '%s'", k)
		}
	}
}

func TestPathTagResolver(t *testing.T) {
	type Handler func() string

	v := struct {
		Root Handler `path:"/"`
		Echo Handler `path:"/Echo"`
	}{}

	prototype, err := structproto.Prototypify(&v,
		&structproto.StructProtoResolveOption{
			TagName:     "path",
			TagResolver: ResolvePathTag,
		})
	if err != nil {
		t.Error(err)
	}

	err = prototype.BindMap(map[string]interface{}{
		"/":     func() string { return "root" },
		"/Echo": func() string { return "echo" },
	}, valuebinder.BuildScalarBinder)
	if err != nil {
		t.Error(err)
	}

	expectRoot := func() string { return "root" }
	if v.Root() != expectRoot() {
		t.Errorf("assert 'Root':: expected '%v', got '%v'", expectRoot(), v.Root())
	}
	expectEcho := func() string { return "echo" }
	if v.Echo() != expectEcho() {
		t.Errorf("assert 'Echo':: expected '%v', got '%v'", expectEcho(), v.Echo())
	}
}

func TestStruct_Map(t *testing.T) {
	type (
		model struct {
			Name        string    `demo:"*NAME"`
			Age         *int      `demo:"*AGE"`
			Alias       []string  `demo:"ALIAS"`
			DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
			Remark      string    `demo:"REMARK,flag1,flag2,,;note the character's personal favor"`
			Numbers     []int     `demo:"NUMBERS"`
		}
	)

	s := model{
		Name:        "luffy",
		Age:         pointy.Int(19),
		Alias:       []string{"lucy"},
		DateOfBirth: time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC),
		Numbers:     []int{5, 12},
	}

	prototype, err := structproto.Prototypify(&s, &structproto.StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		t.Error(err)
	}
	err = prototype.Map(func(field structproto.FieldInfo, rv reflect.Value) error {
		switch rv.Kind() {
		case reflect.String:
			if !rv.IsZero() {
				val := fmt.Sprintf("[%s]", rv.String())
				rv.Set(reflect.ValueOf(val))
			}
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
	expected := model{
		Name:        "[luffy]",
		Age:         pointy.Int(19),
		Alias:       []string{"lucy"},
		DateOfBirth: time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC),
		Numbers:     []int{5, 12},
	}
	if !reflect.DeepEqual(expected, s) {
		t.Errorf("assert:: expected '%+v', got '%+v'", expected, s)
	}
}

var _ structproto.Unmarshaler = new(ExtraInfo)

type ExtraInfo struct {
	Favorite []string
}

func (info *ExtraInfo) UnmarshalStruct(v interface{}) error {
	input, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid ExtraInfo value")
	}
	return json.Unmarshal([]byte(input), &info.Favorite)
}

func TestStruct_WithStructUnmarshaler(t *testing.T) {
	type (
		model struct {
			Name      string    `demo:"*NAME"`
			Age       *int      `demo:"*AGE"`
			ExtraInfo ExtraInfo `demo:"ExtraInfo"`
		}
	)

	s := model{}

	prototype, err := structproto.Prototypify(&s, &structproto.StructProtoResolveOption{
		TagName: "demo",
	})
	err = prototype.BindMap(map[string]interface{}{
		"NAME":      "luffy",
		"AGE":       "19",
		"ExtraInfo": `[ "meat", "king" ]`,
	}, valuebinder.BuildStringBinder)
	if err != nil {
		t.Error(err)
	}

	expected := model{
		Name: "luffy",
		Age:  pointy.Int(19),
		ExtraInfo: ExtraInfo{
			Favorite: []string{
				"meat", "king",
			},
		},
	}
	if !reflect.DeepEqual(expected, s) {
		t.Errorf("assert:: expected '%+v', got '%+v'", expected, s)
	}
}
