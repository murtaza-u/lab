package main

import (
	"fmt"
	"log"
	"time"

	"aidanwoods.dev/go-paseto"
)

const (
	footer   = "/Control/Log"
	implicit = "implicit bytes"
)

func main() {
	t := paseto.NewToken()
	t.SetFooter([]byte(footer))
	t.SetString("role", "daemon")
	t.SetIssuedAt(time.Now())
	t.SetExpiration(time.Now().Add(time.Minute * 5))

	// must be kept secret
	key := paseto.NewV4SymmetricKey()
	fmt.Printf("Secret key: %s\n", key.ExportHex())

	enc := t.V4Encrypt(key, []byte(implicit))
	fmt.Println(enc)

	p := paseto.NewParser()
	token, err := p.ParseV4Local(key, enc, []byte(implicit))
	if err != nil {
		log.Fatal(err)
	}

	role, err := token.GetString("role")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("role: %s\n", role)
	fmt.Printf("footer: %s\n", string(token.Footer()))
}
