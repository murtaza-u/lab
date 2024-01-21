package main

import (
	"fmt"
	"regexp"
)

func main() {
	expr := "^/registry/[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}/daemon$"
	r := regexp.MustCompile(expr)
	cases := []string{
		"/registry/123e4567-e89b-12d3-a456-426655440000/daemon",
		"/registry/c73bcdcc-2669-4bf6-81d3-e4ae73fb11fd/keylog",
		"/registry/C73BCDCC-2669-4Bf6-81d3-E4AE73FB11FD/audio",
		"/registry/c73bcdcc-2669-4bf6-81d3-e4an73fb11fd/daemon",
		"/registry/c73bcdcc26694bf681d3e4ae73fb11fd/daemon",
		"/registry/definitely-not-a-uuid/daemon",
	}
	for _, c := range cases {
		fmt.Printf("case: %s | matched: %v\n", c, r.MatchString(c))
	}
}
