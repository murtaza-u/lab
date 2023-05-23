package main

import (
	"log"

	"github.com/murtaza-u/mtls/client"
)

func main() {
	err := client.Run()
	if err != nil {
		log.Fatal(err)
	}
}
