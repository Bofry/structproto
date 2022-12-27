package valuebinder

import (
	"bytes"
	"encoding/json"
	"net"
	"net/url"
	"reflect"
	"time"

	"github.com/Bofry/structproto/common"
	"github.com/Bofry/types"
)

const (
	errStringValueLength = 24
)

var (
	typeOfString = reflect.TypeOf("")
	typeOfBytes  = reflect.TypeOf([]byte(nil))

	typeOfDuration   = reflect.TypeOf(time.Nanosecond)
	typeOfUrl        = reflect.TypeOf(url.URL{})
	typeOfTime       = reflect.TypeOf(time.Time{})
	typeOfRawContent = reflect.TypeOf(types.RawContent(nil))
	typeOfRawMessage = reflect.TypeOf(json.RawMessage(nil))
	typeOfIP         = reflect.TypeOf(net.IP(nil))
	typeOfBuffer     = reflect.TypeOf(bytes.Buffer{})
)

var _ common.ValueBindProvider = BuildIgnoreBinder

func BuildIgnoreBinder(rv reflect.Value) common.ValueBinder { return nil }
