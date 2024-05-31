package main

import (
	"fmt"
	"log"
	"os"

	"github.com/murtaza-u/arith/client"
	"github.com/murtaza-u/arith/server"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatal("missing command")
	}

	var err error
	cmd := os.Args[1]
	switch cmd {
	case "server":
		err = server.Run()
	case "client":
		err = client.Add(10, 20)
	default:
		err = fmt.Errorf("invalid command: %s", cmd)
	}

	if err != nil {
		log.Fatal(err)
	}
}
