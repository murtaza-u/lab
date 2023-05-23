package main

import "fmt"

func main() {
	var i int
	fmt.Print("Guess a number: ")
	fmt.Scanf("%d", &i)
	fmt.Printf("You guessed: %d\n", i)
}
