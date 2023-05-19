package valuebinder

import (
	"reflect"
	_ "unsafe"

	_ "github.com/cstockton/go-conv"
)

//go:linkname indirectVal github.com/cstockton/go-conv/internal/refutil.IndirectVal
func indirectVal(val reflect.Value) reflect.Value
