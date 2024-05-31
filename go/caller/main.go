package main

import (
	"fmt"
	"os"
)

func main() {
	src := os.Args[0]
	fmt.Printf("src: %s\n", src)
	// abs, err := filepath.Abs(src)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("src: %s\n", abs)

	// newp := "random"
	// rel, err := filepath.Rel(newp, src)
	// fmt.Printf("new path: %s\n", rel)
}
