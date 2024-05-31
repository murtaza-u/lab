package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("SSD: OK")
		time.Sleep(time.Second * 5)
	}
}
