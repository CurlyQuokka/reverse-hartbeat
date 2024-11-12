package certs

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

func CreateClientConfig(ca, crt, key string) (*tls.Config, error) {
	roots, cert, err := loadCerts(ca, crt, key)
	if err != nil {
		return nil, fmt.Errorf("failed to load certs: %w", err)
	}
	return &tls.Config{
		Certificates:       []tls.Certificate{*cert},
		RootCAs:            roots,
		InsecureSkipVerify: true,
	}, nil
}

func CreateServerConfig(ca, crt, key string) (*tls.Config, error) {
	roots, cert, err := loadCerts(ca, crt, key)
	if err != nil {
		return nil, fmt.Errorf("failed to load certs: %w", err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{*cert},
		ClientCAs:    roots,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}, nil
}

func loadCerts(ca, crt, key string) (*x509.CertPool, *tls.Certificate, error) {
	caCertPEM, err := os.ReadFile(ca)
	if err != nil {
		return nil, nil, err
	}

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertPEM)
	if !ok {
		return nil, nil, fmt.Errorf("failed to parse root certificate")
	}

	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		return nil, nil, err
	}

	return roots, &cert, nil
}
