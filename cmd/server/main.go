package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"log"
	"net/http"
	"os"
)

func createServerConfig(ca, crt, key string) (*tls.Config, error) {
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
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    roots,
	}, nil
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello,World!\n"))
}

func main() {
	listen := flag.String("listen", ":4433", "which address to listen")
	ca := flag.String("ca", "./certs/ca.crt", "root certificate")
	crt := flag.String("crt", "./certs/server.crt", "certificate")
	key := flag.String("key", "./certs/server.key", "key")
	flag.Parse()

	config, err := createServerConfig(*ca, *crt, *key)
	if err != nil {
		log.Fatalf("config failed: %s", err.Error())
	}

	server := http.Server{
		Addr:      *listen,
		Handler:   http.HandlerFunc(httpRequestHandler),
		TLSConfig: config,
	}
	defer server.Close()

	server.ListenAndServeTLS("", "")
}
