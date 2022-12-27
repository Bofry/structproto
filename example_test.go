package structproto_test

import (
	"fmt"
	"time"

	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/valuebinder"
)

func ExamplePrototypify() {
	s := struct {
		Name        string    `demo:"*NAME"`
		Age         *int      `demo:"*AGE"`
		Alias       []string  `demo:"ALIAS"`
		DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
		Remark      string    `demo:"REMARK;note the character's personal favor"`
	}{}
	prototype, err := structproto.Prototypify(&s,
		&structproto.StructProtoResolveOption{
			TagName: "demo",
		})
	if err != nil {
		panic(err)
	}

	err = prototype.BindMap(map[string]interface{}{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}, valuebinder.BuildStringBinder)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name       : %q\n", s.Name)
	fmt.Printf("Age        : %d\n", *s.Age)
	fmt.Printf("Alias      : %q\n", s.Alias)
	fmt.Printf("DateOfBirth: %q\n", s.DateOfBirth)
	fmt.Printf("Remark     : %q\n", s.Remark)
	// Output:
	// Name       : "luffy"
	// Age        : 19
	// Alias      : ["lucy"]
	// DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
	// Remark     : ""
}

func ExampleStruct_BindMap() {
	s := struct {
		Name        string    `demo:"*NAME"`
		Age         *int      `demo:"*AGE"`
		Alias       []string  `demo:"ALIAS"`
		DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
		Remark      string    `demo:"REMARK;note the character's personal favor"`
	}{}
	prototype, err := structproto.Prototypify(&s,
		&structproto.StructProtoResolveOption{
			TagName: "demo",
		})
	if err != nil {
		panic(err)
	}

	err = prototype.BindMap(map[string]interface{}{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}, valuebinder.BuildStringBinder)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name       : %q\n", s.Name)
	fmt.Printf("Age        : %d\n", *s.Age)
	fmt.Printf("Alias      : %q\n", s.Alias)
	fmt.Printf("DateOfBirth: %q\n", s.DateOfBirth)
	fmt.Printf("Remark     : %q\n", s.Remark)
	// Output:
	// Name       : "luffy"
	// Age        : 19
	// Alias      : ["lucy"]
	// DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
	// Remark     : ""
}

func ExampleStruct_BindFields() {
	s := struct {
		Name        string    `demo:"*NAME"`
		Age         *int      `demo:"*AGE"`
		Alias       []string  `demo:"ALIAS"`
		DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
		Remark      string    `demo:"REMARK;note the character's personal favor"`
	}{}
	prototype, err := structproto.Prototypify(&s,
		&structproto.StructProtoResolveOption{
			TagName: "demo",
		})
	if err != nil {
		panic(err)
	}

	fieldValues := []structproto.FieldValueEntity{
		{Field: "NAME", Value: "luffy"},
		{Field: "AGE", Value: "19"},
		{Field: "ALIAS", Value: "lucy"},
		{Field: "DATE_OF_BIRTH", Value: "2020-05-05T00:00:00Z"},
	}

	err = prototype.BindFields(fieldValues, valuebinder.BuildStringBinder)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name       : %q\n", s.Name)
	fmt.Printf("Age        : %d\n", *s.Age)
	fmt.Printf("Alias      : %q\n", s.Alias)
	fmt.Printf("DateOfBirth: %q\n", s.DateOfBirth)
	fmt.Printf("Remark     : %q\n", s.Remark)
	// Output:
	// Name       : "luffy"
	// Age        : 19
	// Alias      : ["lucy"]
	// DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
	// Remark     : ""
}

func ExampleFieldValueMap() {
	s := struct {
		Name        string    `demo:"*NAME"`
		Age         *int      `demo:"*AGE"`
		Alias       []string  `demo:"ALIAS"`
		DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
		Remark      string    `demo:"REMARK;note the character's personal favor"`
	}{}
	prototype, err := structproto.Prototypify(&s,
		&structproto.StructProtoResolveOption{
			TagName: "demo",
		})
	if err != nil {
		panic(err)
	}

	err = prototype.BindMap(structproto.FieldValueMap{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}, valuebinder.BuildStringBinder)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name       : %q\n", s.Name)
	fmt.Printf("Age        : %d\n", *s.Age)
	fmt.Printf("Alias      : %q\n", s.Alias)
	fmt.Printf("DateOfBirth: %q\n", s.DateOfBirth)
	fmt.Printf("Remark     : %q\n", s.Remark)
	// Output:
	// Name       : "luffy"
	// Age        : 19
	// Alias      : ["lucy"]
	// DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
	// Remark     : ""
}
