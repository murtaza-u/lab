package heap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type H[T constraints.Ordered] struct {
	slice []T
}

func New[T constraints.Ordered]() H[T] {
	return H[T]{
		slice: make([]T, 0),
	}
}

func (h H[T]) String() string {
	str := "["
	for _, i := range h.slice {
		str += fmt.Sprintf(" %v", i)
	}
	str += " ]"
	return str
}

func (h *H[T]) Len() int {
	return len(h.slice)
}

func (h *H[T]) Insert(data ...T) {
	for _, datum := range data {
		h.slice = append(h.slice, datum)
		h.heapifyUp(len(h.slice) - 1)
	}
}

func (h *H[T]) Extract() T {
	var max T
	if len(h.slice) == 0 {
		return max
	}
	max = h.slice[0]
	len := len(h.slice)
	h.slice[0] = h.slice[len-1]
	h.slice = h.slice[:len-1]
	h.heapifyDown(0)
	return max
}

func (h *H[T]) heapifyDown(idx int) {
	len := len(h.slice)
	left := idx*2 + 1
	right := idx*2 + 2
	if left >= len {
		return
	}

	target := left
	if right < len && h.slice[right] > h.slice[left] {
		target = right
	}

	if h.slice[target] < h.slice[idx] {
		return
	}

	temp := h.slice[idx]
	h.slice[idx] = h.slice[target]
	h.slice[target] = temp
	h.heapifyDown(target)
}

func (h *H[T]) heapifyUp(idx int) {
	parent := (idx - 1) / 2
	if h.slice[parent] < h.slice[idx] {
		temp := h.slice[parent]
		h.slice[parent] = h.slice[idx]
		h.slice[idx] = temp
		h.heapifyUp(parent)
	}
}
