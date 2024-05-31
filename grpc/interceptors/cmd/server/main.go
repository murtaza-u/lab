package main

import (
	"dummy/server"
	"log"
)

func main() {
	err := server.Init()
	if err != nil {
		log.Fatal(err)
	}
}
