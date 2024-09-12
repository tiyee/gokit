package encodes_test

import (
	"encoding/base64"
	"github.com/tiyee/gokit/encodes"
	"github.com/tiyee/gokit/internal/assert"
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
func TestBaseEncode(t *testing.T) {
	as := assert.NewAssert(t, "encodes2")
	for _, s := range []string{"123", "abs", "0"} {
		bu := encodes.Base64URLEncoding([]byte(s))
		dbu, err := encodes.Base64URLDecoding(bu)
		assert.IsNil(as, err)
		assert.Equal(as, s, string(dbu))

		su := encodes.Base64StdEncoding([]byte(s))
		dsu, err := encodes.Base64StdDecoding(su)
		assert.IsNil(as, err)
		assert.Equal(as, s, string(dsu))
	}
}
func TestBase64URLEncoding(t *testing.T) {
	as := assert.NewAssert(t, "TestBase64URLEncoding")
	seg := []byte("111111111111111")
	s1 := encodes.Base64URLEncoding(seg)
	s2 := base64.URLEncoding.EncodeToString(seg)
	assert.Equal(as, s1, s2)
}
func TestBase64StdDecoding(t *testing.T) {
	as := assert.NewAssert(t, "TestBase64StdDecoding")
	s := "eyJleHAiOjE3MjYxMzE1MDMsImlhdCI6MTcyNjA0NTEwMywibmJmIjoxNzI2MDUzNDAwLCJzdWIiOiJkZW1vIiwidXNlcl9pZCI6MX0"
	s1, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(s)
	assert.IsNil(as, err)
	s2, err := encodes.Base64URLDecoding(s)
	assert.IsNil(as, err)
	assert.EqualBytes(as, s1, s2)

}
