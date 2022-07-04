package utils

import (
	"encoding/json"
	"fmt"
)

type number interface {
	int8 | int | int16 | int32 | int64 |
		float32 | float64 |
		uint | uint8 | uint16 | uint32 | uint64
}

type NumberHeap[T number] struct {
	data []T
	comp func(a, b T) bool
}

func NewHeap[T number](comp func(a, b T) bool) *NumberHeap[T] {
	return &NumberHeap[T]{comp: comp}
}

func (h *NumberHeap[T]) Len() int { return len(h.data) }

func (h *NumberHeap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.up(h.Len() - 1)
}

func (h *NumberHeap[T]) Pop() T {
	n := h.Len() - 1
	if n > 0 {
		h.swap(0, n)
		h.down()
	}
	v := h.data[n]
	h.data = h.data[0:n]
	return v
}

func (h *NumberHeap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *NumberHeap[T]) up(jj int) {
	for {
		i := parent(jj)
		if i == jj || h.data[jj] > h.data[i] {
			break
		}
		h.swap(i, jj)
		jj = i
	}
}

func (h *NumberHeap[T]) down() {
	n := h.Len() - 1
	i1 := 0
	for {
		j1 := left(i1)
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		j2 := right(i1)
		if j2 < n && h.comp(h.data[j2], h.data[j1]) {
			j = j2
		}
		if !h.comp(h.data[j], h.data[i1]) {
			break
		}
		h.swap(i1, j)
		i1 = j
	}
}

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return (i * 2) + 1 }
func right(i int) int  { return left(i) + 1 }

func (h *NumberHeap[T]) Dump() []T {
	return h.data
}

func (h NumberHeap[T]) MarshalJSON() ([]byte, error) {
	fmt.Println("Marshalling the Heap")
	j, err := json.Marshal(struct {
		Data []T
	}{
		Data: h.data,
	})
	return j, err
}
