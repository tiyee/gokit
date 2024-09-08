package encodes_test

import (
	"github.com/tiyee/gokit/assert"
	"github.com/tiyee/gokit/encodes"
	"testing"
)

func TestBaseDecode(t *testing.T) {
	is := assert.NewAssert(t, "encodes")

	assert.Equal(is, "GEZGCYTD", encodes.EncodeString("B32Std", "12abc"))
	assert.Equal(is, "12abc", encodes.DecodeString("B32Std", "GEZGCYTD"))

	// b23 hex
	assert.Equal(is, "64P62OJ3", encodes.EncodeString("B32Hex", "12abc"))

	// fmt.Println(time.Now().Format("20060102150405"))
	dateStr := "20230908101122"
	assert.Equal(is, "68O34CPG74O3GC9G64OJ4CG", encodes.EncodeString("B32Hex", dateStr))

	assert.Equal(is, "YWJj", encodes.EncodeString("B64Std", "abc"))

	assert.Equal(is, "MTJhYmM", encodes.EncodeString("B64URL", "12abc"))
}
