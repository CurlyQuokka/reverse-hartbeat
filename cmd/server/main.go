package main

import (
	"flag"
	"log"

	"github.com/curlyquokka/reverse-hearthbeat/pkg/certs"
	"github.com/curlyquokka/reverse-hearthbeat/pkg/server"
)

func main() {
	listen := flag.String("listen", ":4433", "which address to listen")
	ca := flag.String("ca", "./certs/ca.crt", "root certificate")
	crt := flag.String("crt", "./certs/server.crt", "certificate")
	key := flag.String("key", "./certs/server.key", "key")
	flag.Parse()

	config, err := certs.CreateServerConfig(*ca, *crt, *key)
	if err != nil {
		log.Fatalf("config failed: %s", err.Error())
	}

	s := server.New(*listen, config)
	defer s.Close()

	if err := s.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("server failed: %s", err.Error())
	}
}
