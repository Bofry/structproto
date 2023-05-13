package structproto_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/valuebinder"
)

func TestStructBindMap_MissingRequiredField(t *testing.T) {
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

func TestStructBindIterator_WithFieldValueMap(t *testing.T) {
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
	expectedName := "luffy"
	if s.Name != expectedName {
		t.Errorf("assert 'Name':: expected '%v', got '%v'", expectedName, s.Name)
	}
	expectedAge := 19
	if *s.Age != expectedAge {
		t.Errorf("assert 'Age':: expected '%v', got '%v'", expectedAge, s.Age)
	}
	expectedAlias := []string{"lucy"}
	if !reflect.DeepEqual(s.Alias, expectedAlias) {
		t.Errorf("assert 'Alias':: expected '%#v', got '%#v'", expectedAlias, s.Alias)
	}
	expectedDateOfBirth := time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC)
	if s.DateOfBirth != expectedDateOfBirth {
		t.Errorf("assert 'DateOfBirth':: expected '%v', got '%v'", expectedDateOfBirth, s.DateOfBirth)
	}
	expectedNumbers := []int{5, 12}
	if !reflect.DeepEqual(s.Numbers, expectedNumbers) {
		t.Errorf("assert 'Numbers':: expected '%#v', got '%#v'", expectedNumbers, s.Numbers)
	}
}

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

func TestStructBindIterator_Well(t *testing.T) {
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
	expectedName := "luffy"
	if s.Name != expectedName {
		t.Errorf("assert 'Name':: expected '%v', got '%v'", expectedName, s.Name)
	}
	expectedAge := 19
	if *s.Age != expectedAge {
		t.Errorf("assert 'Age':: expected '%v', got '%v'", expectedAge, s.Age)
	}
	expectedAlias := []string{"lucy"}
	if !reflect.DeepEqual(s.Alias, expectedAlias) {
		t.Errorf("assert 'Alias':: expected '%#v', got '%#v'", expectedAlias, s.Alias)
	}
	expectedDateOfBirth := time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC)
	if s.DateOfBirth != expectedDateOfBirth {
		t.Errorf("assert 'DateOfBirth':: expected '%v', got '%v'", expectedDateOfBirth, s.DateOfBirth)
	}
	expectedNumbers := []int{5, 12}
	if !reflect.DeepEqual(s.Numbers, expectedNumbers) {
		t.Errorf("assert 'Numbers':: expected '%#v', got '%#v'", expectedNumbers, s.Numbers)
	}
}

func TestStructBind_MissingRequiredField(t *testing.T) {
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

func TestStructBind_Well(t *testing.T) {
	input := map[string]string{
		"NAME":          "luffy",
		"AGE":           "19",
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
	if err != nil {
		t.Error(err)
	}

	expectedName := "luffy"
	if s.Name != "luffy" {
		t.Errorf("assert 'Name':: expected '%v', got '%v'", expectedName, s.Name)
	}
	expectedAge := 19
	if *s.Age != expectedAge {
		t.Errorf("assert 'Age':: expected '%v', got '%v'", expectedAge, s.Age)
	}
	expectedAlias := []string{"lucy"}
	if !reflect.DeepEqual(s.Alias, expectedAlias) {
		t.Errorf("assert 'Alias':: expected '%#v', got '%#v'", expectedAlias, s.Alias)
	}
	expectedDateOfBirth := time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC)
	if s.DateOfBirth != expectedDateOfBirth {
		t.Errorf("assert 'DateOfBirth':: expected '%v', got '%v'", expectedDateOfBirth, s.DateOfBirth)
	}
}

func TestStructBindMap_Well(t *testing.T) {
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
	err = prototype.BindMap(map[string]interface{}{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}, valuebinder.BuildStringBinder)
	if err != nil {
		t.Error(err)
	}

	if s.Name != "luffy" {
		t.Errorf("assert 'Name':: expected '%v', got '%v'", "luffy", s.Name)
	}
	if *s.Age != 19 {
		t.Errorf("assert 'Age':: expected '%v', got '%v'", 19, s.Age)
	}
	expectedAlias := []string{"lucy"}
	if !reflect.DeepEqual(s.Alias, expectedAlias) {
		t.Errorf("assert 'Alias':: expected '%#v', got '%#v'", expectedAlias, s.Alias)
	}
	expectedDateOfBirth := time.Date(2020, 5, 5, 0, 0, 0, 0, time.UTC)
	if s.DateOfBirth != expectedDateOfBirth {
		t.Errorf("assert 'DateOfBirth':: expected '%v', got '%v'", expectedDateOfBirth, s.DateOfBirth)
	}
}

func TestStructVisitor(t *testing.T) {
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
