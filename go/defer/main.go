package main

import "fmt"

func print(x int) {
	fmt.Println(x)
}

func compute() {
	x := 10
	defer print(x)
	x = 20
	return
}

func main() {
	compute()
}
