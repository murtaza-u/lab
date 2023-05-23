package main

import (
	"log"

	"github.com/murtaza-u/mtls/server"
)

func main() {
	err := server.New()
	if err != nil {
		log.Fatal(err)
	}
}
