package hashmap_test

import (
	"fmt"

	"github.com/murtaza-u/lab/go/data-structures/hashmap"
)

func ExampleEmpty() {
	m := hashmap.New()
	fmt.Print(m)
	// Output: nil
}

func ExampleInsert() {
	m := hashmap.New()
	m.Insert("Randy", 0)
	m.Insert("Erick", 1)
	m.Insert("Mark", 2)
	m.Insert("Spencer", 3)
	m.Insert("Randy", 4)
	fmt.Print(m)
	// Output:
	// Randy(4) -> Spencer(3) -> nil
	// Erick(1) -> nil
	// Mark(2) -> nil
}

func ExampleGet() {
	m := hashmap.New()
	m.Insert("Randy", 0)
	m.Insert("Erick", 1)
	m.Insert("Mark", 2)
	m.Insert("Spencer", 3)
	m.Insert("Randy", 4)
	fmt.Print(m.Get("Spencer"))
	// Output: 3
}

func ExampleDelete() {
	m := hashmap.New()
	m.Insert("Randy", 0)
	m.Insert("Erick", 1)
	m.Insert("Mark", 2)
	m.Insert("Spencer", 3)
	m.Insert("Randy", 4)

	m.Delete("Randy")
	fmt.Print(m)
	// Output:
	// Spencer(3) -> nil
	// Erick(1) -> nil
	// Mark(2) -> nil
}
