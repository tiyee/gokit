package scheduler

// Heap 泛型最小堆实现
type Heap[T any] struct {
	items []T
	less  func(i, j T) bool
}

// NewHeap 创建新的堆
func NewHeap[T any](less func(i, j T) bool) *Heap[T] {
	return &Heap[T]{
		items: make([]T, 0),
		less:  less,
	}
}

// Len 返回堆中元素数量
func (h *Heap[T]) Len() int {
	return len(h.items)
}

// Top 返回堆顶元素（不移除）
func (h *Heap[T]) Top() T {
	if len(h.items) == 0 {
		var zero T
		return zero
	}
	return h.items[0]
}

// Push 向堆中添加元素
func (h *Heap[T]) Push(item T) {
	h.items = append(h.items, item)
	h.up(len(h.items) - 1)
}

// Pop 移除并返回堆顶元素
func (h *Heap[T]) Pop() T {
	if len(h.items) == 0 {
		var zero T
		return zero
	}

	top := h.items[0]
	last := len(h.items) - 1
	h.items[0] = h.items[last]
	h.items = h.items[:last]

	if len(h.items) > 0 {
		h.down(0)
	}

	return top
}

// Remove 移除指定索引的元素
func (h *Heap[T]) Remove(index int) T {
	if index < 0 || index >= len(h.items) {
		var zero T
		return zero
	}

	if index == len(h.items)-1 {
		// 移除最后一个元素
		item := h.items[index]
		h.items = h.items[:index]
		return item
	}

	item := h.items[index]
	last := len(h.items) - 1
	h.items[index] = h.items[last]
	h.items = h.items[:last]

	// 可能需要上浮或下沉
	if index > 0 && h.less(h.items[index], h.items[h.parent(index)]) {
		h.up(index)
	} else {
		h.down(index)
	}

	return item
}

// Update 更新指定索引的元素并重新调整堆
func (h *Heap[T]) Update(index int, item T) {
	if index < 0 || index >= len(h.items) {
		return
	}

	oldItem := h.items[index]
	h.items[index] = item

	// 判断需要上浮还是下沉
	if h.less(item, oldItem) {
		h.up(index)
	} else {
		h.down(index)
	}
}

// up 上浮操作
func (h *Heap[T]) up(index int) {
	for index > 0 {
		parent := h.parent(index)
		if !h.less(h.items[index], h.items[parent]) {
			break
		}
		h.swap(index, parent)
		index = parent
	}
}

// down 下沉操作
func (h *Heap[T]) down(index int) {
	for {
		left := h.leftChild(index)
		right := h.rightChild(index)
		smallest := index

		if left < len(h.items) && h.less(h.items[left], h.items[smallest]) {
			smallest = left
		}

		if right < len(h.items) && h.less(h.items[right], h.items[smallest]) {
			smallest = right
		}

		if smallest == index {
			break
		}

		h.swap(index, smallest)
		index = smallest
	}
}

// parent 返回父节点索引
func (h *Heap[T]) parent(index int) int {
	return (index - 1) / 2
}

// leftChild 返回左子节点索引
func (h *Heap[T]) leftChild(index int) int {
	return 2*index + 1
}

// rightChild 返回右子节点索引
func (h *Heap[T]) rightChild(index int) int {
	return 2*index + 2
}

// swap 交换两个元素
func (h *Heap[T]) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// FindIndex 查找元素索引（线性搜索）
func (h *Heap[T]) FindIndex(predicate func(T) bool) int {
	for i, item := range h.items {
		if predicate(item) {
			return i
		}
	}
	return -1
}
