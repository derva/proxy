package main

import (
	"fmt"
	"net/http"

	"crypto/tls"

	"github.com/derva/proxy/pkg/handlers"
	"github.com/derva/proxy/pkg/logger"
	"github.com/gorilla/mux"
)

func main() {
    conf := handlers.LoadConfFiles()

	l := logger.Logger{
		LogFileName: conf["LOG_FILE_NAME"],
		Location:    conf["LOG_FILE_LOCATION"],
	}

	router := mux.NewRouter()

	tls := &tls.Config{
		Certificates: []tls.Certificate{
			handlers.LoadCertificate(conf["CERTIFICATE_PATH"] + "cert.pem", conf["CERTIFICATE_PATH"]+"key.pem", l),
		},
	}


	router.HandleFunc("/", handlers.HandleWrapper)

	s := http.Server{
		Addr:      ":8080",
		Handler:   router,
		TLSConfig: tls,
	}

	l.Log("Server is starting ...", true)

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}

}
