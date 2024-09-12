package slice_test

import (
	"github.com/tiyee/gokit/internal/assert"
	"github.com/tiyee/gokit/slice"
	"testing"
)

func TestEqual(t *testing.T) {
	as := assert.NewAssert(t, "TestEqual")
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	var c []int = nil
	assert.Equal(as, true, slice.Equal(a, b))
	assert.Equal(as, false, slice.Equal(a, c))
}
func TestKeys(t *testing.T) {
	as := assert.NewAssert(t, "TestKeys")
	m := make(map[int]int)
	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		m[i] = i
		arr = append(arr, i)
	}
	assert.NotEqual(as, true, slice.Equal(arr, slice.Keys(m)))
	assert.Equal(as, true, slice.StrictEqual(arr, slice.Keys(m)))
}
