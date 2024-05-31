package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func hashSalt(payload string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(h)
}

func main() {
	for _, arg := range os.Args[1:] {
		h := hashSalt(arg)
		fmt.Println(h)
	}

	if len(os.Args) >= 2 {
		return
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		h := hashSalt(s.Text())
		fmt.Println(h)
	}

	if err := s.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
