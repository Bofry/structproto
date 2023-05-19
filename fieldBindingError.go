package structproto

import "fmt"

const (
	errStringValueLength = 24
)

type FieldBindingError struct {
	Field string
	Value interface{}
	Err   error
}

func (e *FieldBindingError) Error() string {
	if v, ok := e.Value.(string); ok && len(v) > errStringValueLength {
		return fmt.Sprintf("cannot bind field tag '%s' with value '%s...'. %+v", e.Field, v[:errStringValueLength], e.Err)
	}
	return fmt.Sprintf("cannot bind field tag '%s' with value '%v'. %+v", e.Field, e.Value, e.Err)
}

// Unwrap returns the underlying error.
func (e *FieldBindingError) Unwrap() error {
	return e.Err
}
