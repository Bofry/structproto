structproto - StructPrototype
=============================


## Synopsis

### **Binding struct from map**
```go
import (
  "fmt"
  "time"

  "github.com/Bofry/structproto"
  "github.com/Bofry/structproto/valuebinder"
)


type Character struct {
  Name        string    `demo:"*NAME"`
  Age         *int      `demo:"*AGE"`
  Alias       []string  `demo:"ALIAS"`
  DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
  Remark      string    `demo:"REMARK;note the character's personal favor"`
}

func main() {
  c := Character{}
  prototype, err := structproto.Prototypify(&c,
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
  }, valuebinder.BuildStringArgsBinder)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Name       : %q\n", model.Name)
  fmt.Printf("Age        : %d\n", *model.Age)
  fmt.Printf("Alias      : %q\n", model.Alias)
  fmt.Printf("DateOfBirth: %q\n", model.DateOfBirth)
  fmt.Printf("Remark     : %q\n", model.Remark)
  // Output:
  // Name       : "luffy"
  // Age        : 19
  // Alias      : ["lucy"]
  // DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
  // Remark     : ""
}
```

### **Binding struct from structproto.FieldValueMap**
```go
import (
  "fmt"
  "time"

  "github.com/Bofry/structproto"
  "github.com/Bofry/structproto/valuebinder"
)


type Character struct {
  Name        string    `demo:"*NAME"`
  Age         *int      `demo:"*AGE"`
  Alias       []string  `demo:"ALIAS"`
  DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
  Remark      string    `demo:"REMARK;note the character's personal favor"`
}

func main() {
  c := Character{}
  prototype, err := structproto.Prototypify(&c,
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
  }, valuebinder.BuildStringArgsBinder)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Name       : %q\n", model.Name)
  fmt.Printf("Age        : %d\n", *model.Age)
  fmt.Printf("Alias      : %q\n", model.Alias)
  fmt.Printf("DateOfBirth: %q\n", model.DateOfBirth)
  fmt.Printf("Remark     : %q\n", model.Remark)
  // Output:
  // Name       : "luffy"
  // Age        : 19
  // Alias      : ["lucy"]
  // DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
  // Remark     : ""
}
```

### **Binding struct by FieldValueEntity slice**
```go
import (
  "fmt"
  "time"

  "github.com/Bofry/structproto"
  "github.com/Bofry/structproto/valuebinder"
)


type Character struct {
  Name        string    `demo:"*NAME"`
  Age         *int      `demo:"*AGE"`
  Alias       []string  `demo:"ALIAS"`
  DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
  Remark      string    `demo:"REMARK;note the character's personal favor"`
}

type FieldValue struct {
  Name  string
  Value string
}

func main() {
  c := Character{}
  prototype, err := structproto.Prototypify(&c,
    &structproto.StructProtoResolveOption{
      TagName: "demo",
    })
  if err != nil {
    panic(err)
  }
  
  err = prototype.BindFields([]structproto.FieldValueEntity{
		{Field: "NAME", Value: "luffy"},
		{Field: "AGE", Value: "19"},
		{Field: "ALIAS", Value: "lucy"},
		{Field: "DATE_OF_BIRTH", Value: "2020-05-05T00:00:00Z"},
  }, valuebinder.BuildStringArgsBinder)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Name       : %q\n", model.Name)
  fmt.Printf("Age        : %d\n", *model.Age)
  fmt.Printf("Alias      : %q\n", model.Alias)
  fmt.Printf("DateOfBirth: %q\n", model.DateOfBirth)
  fmt.Printf("Remark     : %q\n", model.Remark)
  // Output:
  // Name       : "luffy"
  // Age        : 19
  // Alias      : ["lucy"]
  // DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
  // Remark     : ""
}
```

### **Binding struct by channel iterator**
```go
import (
  "fmt"
  "time"

  "github.com/Bofry/structproto"
  "github.com/Bofry/structproto/valuebinder"
)


type Character struct {
  Name        string    `demo:"*NAME"`
  Age         *int      `demo:"*AGE"`
  Alias       []string  `demo:"ALIAS"`
  DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
  Remark      string    `demo:"REMARK;note the character's personal favor"`
}

type FieldValue struct {
  Name  string
  Value string
}

func main() {
  c := Character{}
  prototype, err := structproto.Prototypify(&c,
    &structproto.StructProtoResolveOption{
      TagName: "demo",
    })
  if err != nil {
    panic(err)
  }
  
  err = prototype.BindChan(func() <-chan structproto.FieldValueEntity {
    c := make(chan structproto.FieldValueEntity, 1)
    go func() {
      for _, v := range fieldValues {
        c <- v
      }
      close(c)
    }()
    return c
  }(), valuebinder.BuildStringArgsBinder)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Name       : %q\n", model.Name)
  fmt.Printf("Age        : %d\n", *model.Age)
  fmt.Printf("Alias      : %q\n", model.Alias)
  fmt.Printf("DateOfBirth: %q\n", model.DateOfBirth)
  fmt.Printf("Remark     : %q\n", model.Remark)
  // Output:
  // Name       : "luffy"
  // Age        : 19
  // Alias      : ["lucy"]
  // DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
  // Remark     : ""
}
```

### **Binding struct by custom structproto.Iterator**
```go
import (
  "fmt"
  "time"

  "github.com/Bofry/structproto"
  "github.com/Bofry/structproto/valuebinder"
)


type Character struct {
  Name        string    `demo:"*NAME"`
  Age         *int      `demo:"*AGE"`
  Alias       []string  `demo:"ALIAS"`
  DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
  Remark      string    `demo:"REMARK;note the character's personal favor"`
}

var _ Iterator = EntitySet(nil)

type EntitySet [][]string

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

func main() {
  c := Character{}
  prototype, err := structproto.Prototypify(&c,
    &structproto.StructProtoResolveOption{
      TagName: "demo",
    })
  if err != nil {
    panic(err)
  }
  
  err = prototype.BindIterator(EntitySet{
		{"NAME", "luffy"},
		{"AGE", "19"},
		{"ALIAS", "lucy"},
		{"DATE_OF_BIRTH", "2020-05-05T00:00:00Z"},
		{"NUMBERS", "5,12"},
	}, valuebinder.BuildStringArgsBinder)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Name       : %q\n", model.Name)
  fmt.Printf("Age        : %d\n", *model.Age)
  fmt.Printf("Alias      : %q\n", model.Alias)
  fmt.Printf("DateOfBirth: %q\n", model.DateOfBirth)
  fmt.Printf("Remark     : %q\n", model.Remark)
  // Output:
  // Name       : "luffy"
  // Age        : 19
  // Alias      : ["lucy"]
  // DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
  // Remark     : ""
}
```

### **Binding struct by custom StructBinder**
```go
import (
  "fmt"
  "time"

  "github.com/Bofry/structproto"
  "github.com/Bofry/structproto/valuebinder"
)


type Character struct {
  Name        string    `demo:"*NAME"`
  Age         *int      `demo:"*AGE"`
  Alias       []string  `demo:"ALIAS"`
  DateOfBirth time.Time `demo:"DATE_OF_BIRTH;the character's birth of date"`
  Remark      string    `demo:"REMARK;note the character's personal favor"`
}

type MapBinder struct {
  values map[string]string
}

func (b *MapBinder) Init(context *StructProtoContext) error {
  return nil
}

func (b *MapBinder) Bind(field FieldInfo, rv reflect.Value) error {
  name := field.Name()
  if v, ok := b.values[name]; ok {
    return valuebinder.StringArgsBinder(rv).Bind(v)
  }
  return nil
}

func (b *MapBinder) Deinit(context *StructProtoContext) error {
  // validate missing required fields
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

func main() {
  c := Character{}
  prototype, err := structproto.Prototypify(&c,
    &structproto.StructProtoResolveOption{
      TagName: "demo",
    })
  if err != nil {
    panic(err)
  }
  
  err = prototype.Bind(binder, valuebinder.BuildStringArgsBinder)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Name       : %q\n", model.Name)
  fmt.Printf("Age        : %d\n", *model.Age)
  fmt.Printf("Alias      : %q\n", model.Alias)
  fmt.Printf("DateOfBirth: %q\n", model.DateOfBirth)
  fmt.Printf("Remark     : %q\n", model.Remark)
  // Output:
  // Name       : "luffy"
  // Age        : 19
  // Alias      : ["lucy"]
  // DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
  // Remark     : ""
}
```


