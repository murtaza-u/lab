package main

import (
	"encoding/xml"
	"os"

	md "github.com/gomarkdown/markdown"
)

type content struct {
	XMLName xml.Name `xml:"content:encoded"`
	Content string   `xml:",cdata"`
}

func (c *content) set(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	c.Content = string(data)
	return nil
}

func (c *content) toHTML() {
	html := md.ToHTML([]byte(c.Content), nil, nil)
	c.Content = string(html)
}
