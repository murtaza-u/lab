package main

import (
	"dummy/client"
	"log"
)

func main() {
	err := client.Init()
	if err != nil {
		log.Fatal(err)
	}
}
