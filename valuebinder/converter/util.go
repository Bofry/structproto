package converter

import (
	_ "unsafe"

	_ "github.com/cstockton/go-conv"
)

//go:linkname newConvErr github.com/cstockton/go-conv/internal/refconv.newConvErr
func newConvErr(from interface{}, to string) error

//go:linkname indirect github.com/cstockton/go-conv/internal/refutil.Indirect
func indirect(value interface{}) interface{}
