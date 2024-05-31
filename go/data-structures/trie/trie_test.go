package trie_test

import (
	"fmt"

	"github.com/murtaza-u/lab/go/data-structures/trie"
)

func ExampleInsert() {
	trie := trie.New()
	trie.Insert("area")
	trie.Insert("oreo")
	trie.Insert("orange")
	trie.Insert("ant")
	fmt.Print(trie)
	// Output:
	// ant
	// area
	// orange
	// oreo
}

func ExampleSearch() {
	trie := trie.New()
	trie.Insert("area")
	trie.Insert("oreo")
	trie.Insert("orange")
	trie.Insert("ant")
	fmt.Print(trie.Search("oreo"))
	// Output: true
}
