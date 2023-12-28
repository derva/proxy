package handlers

import (
	"crypto/tls"
	"log"

	"github.com/derva/proxy/pkg/logger"
)

func LoadCertificate(certFile, keyFile string, l logger.Logger) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	l.Log("Certificates are loaded", false)
	return cert
}
