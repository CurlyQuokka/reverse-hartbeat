package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/curlyquokka/reverse-hearthbeat/pkg/certs"
)

func main() {
	connect := flag.String("connect", "", "who to connect to")
	ca := flag.String("ca", "./certs/ca.crt", "root certificate")
	crt := flag.String("crt", "./certs/client.crt", "certificate")
	key := flag.String("key", "./certs/client.key", "key")
	flag.Parse()

	if *connect == "" {
		log.Fatalf("please specify connection address")
	}

	config, err := certs.CreateClientConfig(*ca, *crt, *key)
	if err != nil {
		log.Fatalf("config failed: %s", err.Error())
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}
	client := &http.Client{Transport: tr}

	body := []byte(`{
		"id": "some-service",
		"value": "running"
	}`)

	resp, err := client.Post(*connect, "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	msg, _ := io.ReadAll(resp.Body)
	fmt.Print(string(msg))
	return
}
