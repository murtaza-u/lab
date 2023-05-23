package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var Home string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	Home = home
}

func main() {
	var err error

	pwd, err = filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		inp := s.Text()
		inp = strings.TrimSpace(inp)
		if inp == "" {
			continue
		}

		out, err := run(inp)
		if err != nil {
			fmt.Printf("err: %s\n", err.Error())
			continue
		}

		fmt.Println(out)
	}

	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

var pwd string

func run(inp string) (string, error) {
	s := strings.Split(inp, " ")
	cmd := s[0]

	var args string

	if len(s) >= 2 {
		args = strings.Join(s[1:], " ")
	}

	switch cmd {
	case "pwd":
		return pwd, nil
	case "ls", "list":
		list, err := ls(args)
		if err != nil {
			return "", err
		}

		return strings.Join(list, "\n"), nil
	case "cd":
		if args == "" {
			args = Home
		}

		info, err := os.Stat(args)
		if err != nil {
			args = filepath.Join(pwd, args)
			info, err = os.Stat(args)
			if err != nil {
				return "", err
			}
		}

		if !info.IsDir() {
			return "", fmt.Errorf("%s is not a directory", args)
		}

		abs, err := filepath.Abs(args)
		if err != nil {
			return "", err
		}

		pwd = abs
		return pwd, err
	case "get":
	case "put":
	case "rm", "del", "delete", "remove":
	default:
		return "", fmt.Errorf("invalid command")
	}

	return "", nil
}

func ls(dir string) ([]string, error) {
	var list []string

	if dir == "" {
		dir = pwd
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		dir = filepath.Join(pwd, dir)
		entries, err = os.ReadDir(dir)
		if err != nil {
			return nil, err
		}
	}

	for _, e := range entries {
		list = append(list, e.Name())
	}

	return list, nil
}
