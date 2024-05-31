package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	baseURL := "https://account.murtazau.xyz"
	baseURL = strings.TrimSpace(baseURL)
	baseURL = strings.TrimSuffix(baseURL, "/")

	errName := url.QueryEscape("invalid_request")
	errDesc := url.QueryEscape("missing scope")

	uri := fmt.Sprintf(
		"%s?error=%s&error_description=%s",
		baseURL, errName, errDesc)

	url, err := url.ParseRequestURI(uri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
}
