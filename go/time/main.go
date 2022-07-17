package main

import (
	"fmt"
	"log"
	"time"
)

const (
	format = "20060102150405" // YYYYMMDDhhmmss
	isosec = "20220717083059"
)

func main() {
	t, err := time.Parse(format, isosec)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("UTC: %v\nLocal: %v\n", t, t.Local())
}
