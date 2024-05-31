package main

import (
	"errors"
	"flag"
)

const (
	Author    = "Murtaza Udaipurwala"
	Email     = "murtaza@murtazau.xyz"
	Copyright = "Creative Commons Zero v1.0 Universal"
	Desc      = "Updates from Murtaza Udaipurwala"
	Link      = "https://murtazau.xyz"
	Lang      = "en"
)

var ErrMissingArguments = errors.New("missing arguments")

type cmd struct {
	title    string
	link     string
	isosec   string
	file     string
	template bool
}

func (c *cmd) parse() error {
	flag.StringVar(&c.title, "title", "", "title")
	flag.StringVar(&c.link, "link", "", "link to article")
	flag.StringVar(&c.isosec, "isosec", "", "isosec identifier")
	flag.StringVar(&c.file, "file", "", "path to article file(md)")
	flag.BoolVar(&c.template, "template", false, "generate template")

	flag.Parse()

	if c.template {
		return nil
	}

	if flag.NFlag() < 4 {
		return ErrMissingArguments
	}

	return nil
}
