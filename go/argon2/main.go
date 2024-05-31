package main

import (
	"fmt"
	"log"

	"github.com/alexedwards/argon2id"
)

const secret = "nEwbLhBNwcwV1kjvzQlm9Mu8AblTQhWFG6DHOXEEJ0v5Rr5BEwi3703k1sTlaeHgs0dDsa"

func main() {
	hash, err := argon2id.CreateHash(secret, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hash)
}
