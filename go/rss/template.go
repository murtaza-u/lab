package main

import (
	"encoding/xml"
	"fmt"
)

const (
	RSSVer       = "2.0"
	RSSNamespace = "http://purl.org/rss/1.0/modules/content/"
)

type template struct {
	t                string
	XMLName          xml.Name `xml:"rss"`
	Version          string   `xml:"version,attr"`
	ContentNamespace string   `xml:"xmlns:content,attr"`
	Channel          *channel
}

type channel struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`
	Desc           string   `xml:"description"`
	Link           string   `xml:"link"`
	Lang           string   `xml:"language"`
	Copyright      string   `xml:"copyright"`
	ManagingEditor string   `xml:"managingEditor"`
	Placeholder    string   `xml:",comment"`
}

func (t *template) toXML() ([]byte, error) {
	buf, err := xml.MarshalIndent(t, "", "  ")
	return buf, err
}

func (t *template) gen() error {
	t.Version = RSSVer
	t.ContentNamespace = RSSNamespace
	t.Channel = new(channel)
	t.Channel.Title = Author
	t.Channel.Desc = Desc
	t.Channel.Link = Link
	t.Channel.Lang = Lang
	t.Channel.Copyright = Copyright
	t.Channel.ManagingEditor = fmt.Sprintf("%s (%s)", Email, Author)
	t.Channel.Placeholder = "place item here"

	buf, err := t.toXML()
	if err != nil {
		return err
	}

	t.t = xml.Header + string(buf)
	return nil
}

func (t *template) String() string {
	return t.t
}
