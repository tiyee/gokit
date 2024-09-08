package collection

import (
	"cmp"
	"fmt"
)

// MaxHeap 是一个泛型最大堆结构
type MaxHeap[T cmp.Ordered] struct {
	data []T
}

// NewMaxHeap 创建一个新的最大堆
func NewMaxHeap[T cmp.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{
		data: make([]T, 0),
	}
}

// Push 向堆中插入一个元素
func (h *MaxHeap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.heapifyUp(len(h.data) - 1)
}

// Pop 从堆中移除并返回最大元素
func (h *MaxHeap[T]) Pop() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}

	max_ := h.data[0]
	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]
	h.heapifyDown(0)

	return max_, true
}

// heapifyUp 从下往上调整堆
func (h *MaxHeap[T]) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.data[index] > h.data[parentIndex] {
			h.data[index], h.data[parentIndex] = h.data[parentIndex], h.data[index]
			index = parentIndex
		} else {
			break
		}
	}
}

// heapifyDown 从上往下调整堆
func (h *MaxHeap[T]) heapifyDown(index int) {
	length := len(h.data)
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		largestIndex := index

		if leftChildIndex < length && h.data[largestIndex] < h.data[leftChildIndex] {
			largestIndex = leftChildIndex
		}

		if rightChildIndex < length && h.data[largestIndex] < h.data[rightChildIndex] {
			largestIndex = rightChildIndex
		}

		if largestIndex != index {
			h.data[index], h.data[largestIndex] = h.data[largestIndex], h.data[index]
			index = largestIndex
		} else {
			break
		}
	}
}

func main() {
	// 创建一个最大堆，比较函数为 less
	heap := NewMaxHeap[int]()

	// 插入一些元素
	heap.Push(3)
	heap.Push(1)
	heap.Push(4)
	heap.Push(1)
	heap.Push(5)
	heap.Push(9)
	heap.Push(2)
	heap.Push(6)

	// 弹出并打印最大元素
	for {
		if val, ok := heap.Pop(); ok {
			fmt.Println(val)
		} else {
			break
		}
	}
}
