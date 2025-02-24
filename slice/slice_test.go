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
func TestMerge(t *testing.T) {
	as := assert.NewAssert(t, "TestMerge")
	a := []int{1, 2, 3}
	b := []int{1, 2, 3, 4}
	c := []int{5, 2, 3}
	d := []int{1, 2, 3, 1, 2, 3, 4, 5, 2, 3}
	r := slice.Merge(a, b, c)
	assert.Equal(as, true, slice.StrictEqual(d, r))
}
func TestFirstOr(t *testing.T) {
	as := assert.NewAssert(t, "TestFirstOr")
	a := []int{1, 2, 3}
	b := []int{}
	assert.Equal(as, 1, slice.FirstOr(a, 9))
	assert.Equal(as, 9, slice.FirstOr(b, 9))

}
func TestLastOr(t *testing.T) {
	as := assert.NewAssert(t, "TestLastOr")
	a := []int{1, 2, 3}
	b := []int{}
	assert.Equal(as, 3, slice.LastOr(a, 9))
	assert.Equal(as, 9, slice.LastOr(b, 9))

}

func TestReverse(t *testing.T) {
	as := assert.NewAssert(t, "TestLastOr")
	a := []int{1, 2, 3}
	b := []int{3, 2, 1}
	slice.Reverse(b)
	for i := 0; i < 3; i++ {
		assert.Equal(as, b[i], a[i])
	}
}