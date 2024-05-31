package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func readPrivateKeyFromFile(fname string) (*rsa.PrivateKey, error) {
	// read the key file
	keyData, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	// decode PEM data
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// parse the private key
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func readPublicKeyFromFile(fname string) (*rsa.PublicKey, error) {
	// read the key file
	keyData, err := os.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	// decode PEM data
	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}

	// parse the public key
	ifc, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to parse public key")
	}

	return key, nil
}

func main() {
	privateKeyFile := "private_key.pem"
	publicKeyFile := "public_key.pem"

	// read private key
	privateKey, err := readPrivateKeyFromFile(privateKeyFile)
	if err != nil {
		fmt.Println("Error reading private key:", err)
		return
	}
	fmt.Println("Private Key:", privateKey)

	// read public key
	publicKey, err := readPublicKeyFromFile(publicKeyFile)
	if err != nil {
		fmt.Println("Error reading public key:", err)
		return
	}
	fmt.Println("Public Key:", publicKey)
}
