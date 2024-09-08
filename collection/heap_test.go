package collection_test

import (
	"github.com/tiyee/gokit/assert"
	"github.com/tiyee/gokit/collection"
	"sort"
	"testing"
)

func TestHeap(t *testing.T) {
	as := assert.NewAssert(t, "TestHeap")
	inputs := []int{2, 3, 4, 1, 2, 9}
	heap := collection.NewMaxHeap[int]()
	for _, input := range inputs {
		heap.Push(input)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inputs)))
	for _, input := range inputs {
		if n, ok := heap.Pop(); ok {
			assert.Equal(as, input, n)
		}
	}
}
