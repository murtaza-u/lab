package main

import (
	"encoding/xml"
	"time"
)

type item struct {
	i       string
	XMLName xml.Name  `xml:"item"`
	Title   string    `xml:"title"`
	Link    string    `xml:"link"`
	Author  string    `xml:"author,omitempty"`
	Guid    string    `xml:"guid,omitempty"`
	PubDate time.Time `xml:"pubDate,omitempty"`
	Content *content
}

func (i *item) toXML() error {
	buf, err := xml.MarshalIndent(i, "    ", "  ")
	if err != nil {
		return err
	}

	i.i = string(buf)
	return nil
}

func (i *item) String() string {
	return i.i
}
