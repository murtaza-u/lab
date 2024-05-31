package heap_test

import (
	"fmt"

	"github.com/murtaza-u/lab/go/data-structures/heap"
)

func ExampleInsert() {
	h := heap.New[int]()
	h.Insert(10, -1, 20, 30, 0, 50, 5, 5, 4, 3)

	fmt.Print(h)
	// Output: [ 50 20 30 5 3 10 5 -1 4 0 ]
}

func ExampleExtract() {
	h := heap.New[int]()
	h.Insert(10, -1, 20, 30, 0, 50, 5, 5, 4, 3)

	for h.Len() != 0 {
		max := h.Extract()
		fmt.Printf("%d ", max)
	}
	// Output: 50 30 20 10 5 5 4 3 0 -1
}
