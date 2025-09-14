# structproto - StructPrototype

[![Go Reference](https://pkg.go.dev/badge/github.com/Bofry/structproto.svg)](https://pkg.go.dev/github.com/Bofry/structproto)
[![Go Report Card](https://goreportcard.com/badge/github.com/Bofry/structproto)](https://goreportcard.com/report/github.com/Bofry/structproto)

A high-performance Go library for binding data to structs using reflection with configurable field mapping and type conversion.

## Features

- **High Performance**: Optimized reflection usage with pre-allocated maps and efficient search algorithms
- **Flexible Binding**: Support for maps, slices, channels, and custom iterators
- **Configurable Tags**: Custom struct tags for field mapping and validation
- **Type Safety**: Built-in type conversion with comprehensive error handling
- **Multiple Sources**: Bind from various data sources with unified interface
- **Required Fields**: Built-in validation for required field checking

## Installation

```bash
go get github.com/Bofry/structproto
```

## Quick Start

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
  fmt.Printf("Name       : %q\n", c.Name)
  fmt.Printf("Age        : %d\n", *c.Age)
  fmt.Printf("Alias      : %q\n", c.Alias)
  fmt.Printf("DateOfBirth: %q\n", c.DateOfBirth)
  fmt.Printf("Remark     : %q\n", c.Remark)
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
  fmt.Printf("Name       : %q\n", c.Name)
  fmt.Printf("Age        : %d\n", *c.Age)
  fmt.Printf("Alias      : %q\n", c.Alias)
  fmt.Printf("DateOfBirth: %q\n", c.DateOfBirth)
  fmt.Printf("Remark     : %q\n", c.Remark)
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
  fmt.Printf("Name       : %q\n", c.Name)
  fmt.Printf("Age        : %d\n", *c.Age)
  fmt.Printf("Alias      : %q\n", c.Alias)
  fmt.Printf("DateOfBirth: %q\n", c.DateOfBirth)
  fmt.Printf("Remark     : %q\n", c.Remark)
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
  fmt.Printf("Name       : %q\n", c.Name)
  fmt.Printf("Age        : %d\n", *c.Age)
  fmt.Printf("Alias      : %q\n", c.Alias)
  fmt.Printf("DateOfBirth: %q\n", c.DateOfBirth)
  fmt.Printf("Remark     : %q\n", c.Remark)
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

var _ structproto.Iterator = EntitySet(nil)

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
  fmt.Printf("Name       : %q\n", c.Name)
  fmt.Printf("Age        : %d\n", *c.Age)
  fmt.Printf("Alias      : %q\n", c.Alias)
  fmt.Printf("DateOfBirth: %q\n", c.DateOfBirth)
  fmt.Printf("Remark     : %q\n", c.Remark)
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
  
  binder := &MapBinder{
    values: map[string]string{
      "NAME":          "luffy",
      "AGE":           "19",
      "ALIAS":         "lucy",
      "DATE_OF_BIRTH": "2020-05-05T00:00:00Z",
    },
  }

  err = prototype.Bind(binder)
  if err != nil {
    panic(err)
  }
  fmt.Printf("Name       : %q\n", c.Name)
  fmt.Printf("Age        : %d\n", *c.Age)
  fmt.Printf("Alias      : %q\n", c.Alias)
  fmt.Printf("DateOfBirth: %q\n", c.DateOfBirth)
  fmt.Printf("Remark     : %q\n", c.Remark)
  // Output:
  // Name       : "luffy"
  // Age        : 19
  // Alias      : ["lucy"]
  // DateOfBirth: "2020-05-05 00:00:00 +0000 UTC"
  // Remark     : ""
}
```

## API Reference

### Core Types

#### StructProtoResolver

The main resolver for creating struct prototypes.

```go
type StructProtoResolver struct {
    // Configuration fields
}

func NewStructProtoResolver(option *StructProtoResolveOption) *StructProtoResolver
func (r *StructProtoResolver) Resolve(target interface{}) (*Struct, error)
```

#### Struct

Represents a resolved struct prototype with binding capabilities.

```go
type Struct struct {
    // Internal fields
}

// Binding methods
func (s *Struct) BindMap(values map[string]interface{}, buildValueBinder ValueBindProvider) error
func (s *Struct) BindFields(values []FieldValueEntity, buildValueBinder ValueBindProvider) error
func (s *Struct) BindChan(iterator <-chan FieldValueEntity, buildValueBinder ValueBindProvider) error
func (s *Struct) BindIterator(iterator Iterator, buildValueBinder ValueBindProvider) error
func (s *Struct) Bind(binder StructBinder) error
func (s *Struct) Map(mapper StructMapper) error
func (s *Struct) Visit(visitor StructVisitor)
```

### Configuration Options

```go
type StructProtoResolveOption struct {
    TagName             string      // Custom tag name for field mapping
    TagResolver         TagResolver // Custom tag resolution logic
    CheckDuplicateNames bool        // Enable duplicate field name checking
}
```

### Tag Syntax

Fields can be configured using struct tags:

```go
type Example struct {
    RequiredField string `demo:"*FIELD_NAME"`           // Required field
    OptionalField string `demo:"FIELD_NAME"`            // Optional field
    WithDesc      string `demo:"FIELD;description here"` // Field with description
}
```

- `*` prefix marks required fields
- `;` separates field name from description

## Performance

This library has been optimized for high-performance scenarios:

- **Pre-allocated Maps**: Reduces memory allocations during struct initialization
- **Efficient Search**: Optimized binary search for field lookups
- **Minimal Goroutines**: Eliminates unnecessary concurrency overhead
- **Reflection Caching**: Reuses reflection metadata where possible

### Benchmarks

```go
// Run benchmarks
go test -bench=. -benchmem
```

## Requirements

- Go 1.21 or later
- No external dependencies (except for testing)

## Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific package tests
go test ./valuebinder -v
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests for your changes
5. Ensure all tests pass (`go test ./...`)
6. Run linting (`go fmt ./... && go vet ./...`)
7. Commit your changes (`git commit -am 'Add amazing feature'`)
8. Push to the branch (`git push origin feature/amazing-feature`)
9. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Examples

For more comprehensive examples, check out the [example_test.go](example_test.go) file.
