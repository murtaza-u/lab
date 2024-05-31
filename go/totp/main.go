package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	secret := os.Args[1]
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 8)
	interval := time.Now().Unix() / 30
	binary.BigEndian.PutUint64(buf, uint64(interval))

	hash := hmac.New(sha1.New, key)
	hash.Write(buf)
	h := hash.Sum(nil)

	o := h[19] & 15
	var header uint32
	r := bytes.NewReader(h[o : o+4])
	err = binary.Read(r, binary.BigEndian, &header)
	if err != nil {
		log.Fatal(err)
	}

	h12 := (int(header) & 0x7fffffff) % 100000
	otp := strconv.Itoa(int(h12))
	fmt.Println(otp)
}
