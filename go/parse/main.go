package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parse() error {
	args := os.Args[1:]
	if len(args) == 0 {
		return fmt.Errorf("missing argument")
	}

	var id, title, icon string
	var body []string

	for _, arg := range args {
		splits := strings.Split(arg, "=")
		if len(splits) < 2 {
			body = append(body, arg)
			continue
		}

		switch splits[0] {
		case "--id":
			id = strings.Join(splits[1:], "=")
		case "--title":
			title = strings.Join(splits[1:], "=")
		case "--icon":
			icon = strings.Join(splits[1:], "=")
		default:
			body = append(body, arg)
		}
	}

	fmt.Printf(
		"id: %s | title: %s | icon: %s | body: %q\n",
		id, title, icon, strings.Join(body, " "),
	)

	return nil
}

func main() {
	err := parse()
	if err != nil {
		log.Fatal(err)
	}
}
