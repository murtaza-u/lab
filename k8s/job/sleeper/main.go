package main

import (
	"log"
	"time"
)

func main() {
	log.Println("sleeping for 2 min")
	time.Sleep(time.Minute * 2)
	log.Println("sleep completed")
}
