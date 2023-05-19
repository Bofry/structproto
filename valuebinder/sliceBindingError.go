package valuebinder

import "fmt"

type SliceBindingError struct {
	Value interface{}
	Kind  string
	Index int
	Err   error
}

func (e *SliceBindingError) Error() string {
	if v, ok := e.Value.(string); ok {
		if len(v) > errStringValueLength {
			return fmt.Sprintf("cannot bind type %s with value (type %[1]T) '%v' at %d. %+v", e.Kind, v[:errStringValueLength], e.Index, e.Err)
		}
	}
	return fmt.Sprintf("cannot bind type %s with value (type %[1]T) '%v'at %d. %+v", e.Kind, e.Value, e.Index, e.Err)
}

// Unwrap returns the underlying error.
func (e *SliceBindingError) Unwrap() error {
	return e.Err
}
