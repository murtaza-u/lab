package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Run() error {
	cert, err := tls.LoadX509KeyPair("certs/iot-cert.pem", "certs/iot-key.pem")
	if err != nil {
		return err
	}

	caCert, err := os.ReadFile("certs/ca-cert.pem")
	if err != nil {
		return err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	r, err := client.Get("https://api.murtazau.xyz/")
	if err != nil {
		return err
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", body)
	return nil
}
