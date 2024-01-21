package main

type Request struct {
	Body   []byte
	Method string
	URL    string
	Header Header
}

type Header map[string][]string

type Response struct {
	Data   []byte
	Status int
}
