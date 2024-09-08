package hashlib_test

import (
	"github.com/tiyee/gokit/hashlib"
	"github.com/tiyee/gokit/internal/assert"
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		src  []byte
		algo hashlib.HashType
		want string
	}{
		{[]byte("abc12"), "crc32", "b744b523"},
		{[]byte("abc12"), "crc64", "41b31776c4200000"},
		{[]byte("abc12"), "md5", "b2157e7b2ae716a747597717f1efb7a0"},
		{[]byte("abc12"), "sha1", "8fe670fef2b8c74ef8987cdfccdb32e96ad4f9a2"},
	}
	is := assert.NewAssert(t, "hashlib")
	for _, tt := range tests {
		assert.Equal(is, tt.want, hashlib.Hash(tt.algo, tt.src))
	}

	assert.Panics(is, func() {
		hashlib.Hash("unknown", nil)
	})
}

func TestHash32(t *testing.T) {
	tests := []struct {
		src  []byte
		algo hashlib.HashType
		want string
	}{
		{[]byte("abc12"), "crc32", "MT2BA8O"},
		{[]byte("abc12"), "crc64", "86PHETM440000"},
		{[]byte("abc12"), "md5", "M8ANSUPASSBAEHQPESBV3RTNK0"},
		{[]byte("abc12"), "sha1", "HVJ71VNIN33KTU4OFJFSPMPIT5LD9UD2"},
	}
	is := assert.NewAssert(t, "hashlib.32")
	for _, tt := range tests {
		assert.Equal(is, tt.want, hashlib.Hash32(tt.algo, tt.src))
	}
}

func TestHash64(t *testing.T) {
	tests := []struct {
		src  []byte
		algo hashlib.HashType
		want string
	}{
		{[]byte("abc12"), "crc32", "t0S1Iw"},
		{[]byte("abc12"), "crc64", "QbMXdsQgAAA"},
		{[]byte("abc12"), "md5", "shV+eyrnFqdHWXcX8e+3oA"},
		{[]byte("abc12"), "sha1", "j+Zw/vK4x074mHzfzNsy6WrU+aI"},
	}
	is := assert.NewAssert(t, "hashlib.64")
	for _, tt := range tests {
		assert.Equal(is, tt.want, hashlib.Hash64(tt.algo, tt.src))
	}
}

func TestMd5(t *testing.T) {
	is := assert.NewAssert(t, "hashlib.md5")
	assert.NotEqual(is, "", hashlib.MD5([]byte("abc")))
	assert.NotEqual(is, "", hashlib.MD5([]byte{12, 34}))

	assert.Equal(is, "202cb962ac59075b964b07152d234b70", hashlib.MD5([]byte("123")))

	// short md5
	assert.Equal(is, "ac59075b964b0715", hashlib.ShortMD5([]byte("123")))
	assert.Equal(is, "3cd24fb0d6963f7d", hashlib.ShortMD5([]byte("abc")))
}
