package valuebinder

import (
	"bytes"
	"encoding/json"
	"net"
	"net/url"
	"reflect"
	"time"

	"github.com/Bofry/structproto/internal"
	"github.com/Bofry/types"
)

var (
	typeOfDuration   = reflect.TypeOf(time.Nanosecond)
	typeOfUrl        = reflect.TypeOf(url.URL{})
	typeOfTime       = reflect.TypeOf(time.Time{})
	typeOfRawContent = reflect.TypeOf(types.RawContent(nil))
	typeOfRawMessage = reflect.TypeOf(json.RawMessage(nil))
	typeOfIP         = reflect.TypeOf(net.IP(nil))
	typeOfBuffer     = reflect.TypeOf(bytes.Buffer{})
)

var _ internal.ValueBindProvider = BuildIgnoreBinder

func BuildIgnoreBinder(rv reflect.Value) internal.ValueBinder { return nil }
