package main

import (
	"fmt"
	"log"

	"github.com/jaypipes/ghw"
)

func main() {
	mem, err := ghw.Memory()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mem)

	phys := mem.TotalPhysicalBytes
	usable := mem.TotalUsableBytes
	fmt.Printf("The bootloader consumes %d bytes of RAM\n", phys-usable)
}
