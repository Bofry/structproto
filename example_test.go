package structproto_test

import (
	"time"

	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/valuebinder"
)

func Example() {
	model := struct {
		Name       string    `demo:"*NAME"`
		Age        *int      `demo:"*AGE"`
		Alias      []string  `demo:"ALIAS"`
		DatOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
		Remark     string    `demo:"REMARK;note the character's personal favor"`
	}{}
	prototype, err := structproto.Prototypify(&model, &structproto.StructProtoResolveOption{
		TagName: "demo",
	})
	if err != nil {
		panic(err)
	}

	err = prototype.BindFields(map[string]interface{}{
		"NAME":          "luffy",
		"AGE":           "19",
		"ALIAS":         "lucy",
		"DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
	}, valuebinder.BuildStringArgsBinder)
	if err != nil {
		panic(err)
	}

	
}
