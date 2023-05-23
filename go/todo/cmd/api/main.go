package main

import (
	"log"

	"github.com/murtaza-u/todo"
)

func main() {
	err := api.NewServer().Run()
	if err != nil {
		log.Fatal(err)
	}
}
