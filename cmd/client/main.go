package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func createClientConfig(ca, crt, key string) (*tls.Config, error) {
	caCertPEM, err := os.ReadFile(ca)
	if err != nil {
		return nil, err
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertPEM)
	if !ok {
		panic("failed to parse root certificate")
	}

	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            roots,
		InsecureSkipVerify: true,
	}, nil
}

func main() {
	connect := flag.String("connect", "", "who to connect to")
	ca := flag.String("ca", "./certs/ca.crt", "root certificate")
	crt := flag.String("crt", "./certs/client.crt", "certificate")
	key := flag.String("key", "./certs/client.key", "key")
	flag.Parse()

	if *connect == "" {
		log.Fatalf("please specify connection address")
	}

	config, err := createClientConfig(*ca, *crt, *key)
	if err != nil {
		log.Fatalf("config failed: %s", err.Error())
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(*connect)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	msg, _ := io.ReadAll(resp.Body)
	fmt.Print(string(msg))
	return
}
