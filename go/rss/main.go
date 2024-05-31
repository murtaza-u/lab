package main

import (
	"fmt"
	"os"
)

func out(err error) {
	if err == nil {
		return
	}

	fmt.Printf("error: %s\n", err.Error())
	os.Exit(1)
}

func load(i *item, c *cmd) error {
	i.Title = c.title
	i.Link = c.link
	i.Author = fmt.Sprintf("%s (%s)", Email, Author)
	i.Guid = c.isosec

	t, err := parseIsosec(c.isosec)
	if err != nil {
		return err
	}
	i.PubDate = t

	i.Content = new(content)
	err = i.Content.set(c.file)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	cmd := new(cmd)
	err := cmd.parse()
	out(err)

	if cmd.template {
		t := new(template)
		err := t.gen()
		out(err)
		fmt.Println(t)
		return
	}

	i := new(item)
	err = load(i, cmd)
	out(err)

	i.Content.toHTML()

	err = i.toXML()
	out(err)
	fmt.Println(i)
}
