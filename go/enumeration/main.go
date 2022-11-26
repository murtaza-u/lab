package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	hostname, _ := os.Hostname()
	count, _ := cpu.Counts(false)
	mem, _ := mem.VirtualMemory()
	user, _ := user.Current()
	fmt.Println(hostname, count, mem.Total, user.Name)
}
