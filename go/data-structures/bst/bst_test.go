package bst_test

import (
	"fmt"

	"github.com/murtaza-u/lab/go/data-structures/bst"
)

func ExampleInsert() {
	bst := bst.New[int]()
	bst.Insert(50, 10, 20, 60)
	fmt.Print(bst)
	// Output:
	// 50 -> (10, 60)
	// 10 -> (nil, 20)
	// 20 -> (nil, nil)
	// 60 -> (nil, nil)
}

func ExampleSearch() {
	bst := bst.New[int]()
	bst.Insert(50, 10, 20, 60, 0)
	fmt.Print(bst.Search(0))
	// Output: true
}
