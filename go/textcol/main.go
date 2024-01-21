package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/acarl005/textcol"
)

func getTermWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	numsStr := strings.Trim(string(out), "\n ")
	width, _ := strconv.Atoi(strings.Split(numsStr, " ")[1])
	return width
}

func main() {
	fmt.Println(getTermWidth())

	items := []string{
		"hello",
		"world",
		"foo",
		"bar",
		"blah",
	}
	// pass pointer to array of strings and a margin value. this will ensure at least 4 spaces appear to the right of each cell
	textcol.PrintColumns(&items, 4)
}
