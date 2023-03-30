package helps

import (
	"encoding/json"
	"github.com/tiyee/gokit/pkg/constraints"
	"strconv"
)

func ToBool[T constraints.Integer](i T) bool {
	var j T = 0
	return i != j
}

func test() bool {
	return ToBool(1)
}

type Cmp interface {
	Less()
}

func BytesToInt64(bs []byte, defaultValue int64) int64 {
	if len(bs) == 0 {
		return defaultValue
	}
	var ret int64
	if n, err := strconv.ParseInt(string(bs), 10, 64); err == nil {
		ret = n
	} else {
		ret = defaultValue
	}
	return ret
}
func t() {
	var x struct{ A int64 }
	if err := json.Unmarshal([]byte(""), &x); err == nil {

	}
}
