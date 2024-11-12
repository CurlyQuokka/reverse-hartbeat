package server

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/curlyquokka/reverse-hearthbeat/pkg/status"
)

func New(listen string, config *tls.Config) *http.Server {
	sm := http.NewServeMux()
	sm.HandleFunc("/add", httpRequestHandler)

	return &http.Server{
		Addr:      listen,
		Handler:   sm,
		TLSConfig: config,
	}
}

func httpRequestHandler(w http.ResponseWriter, req *http.Request) {
	s := &status.Info{}
	err := json.NewDecoder(req.Body).Decode(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got: %v\n", s)
	w.Write([]byte("Hello,World!\n"))
}
