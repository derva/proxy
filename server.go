package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/derva/proxy/pkg/conf"
	"github.com/derva/proxy/pkg/handlers"
	"github.com/derva/proxy/pkg/logger"
	"github.com/gorilla/mux"
)

func main() {
	l := logger.Logger{
		LogFileName: "proxy.log",
		Location:    "/var/log/",
	}

	conf := conf.LoadConfFiles()
	router := mux.NewRouter()

	tls := &tls.Config{
		Certificates: []tls.Certificate{
			handlers.LoadCertificate(conf["certificate"]+"cert.pem", conf["certificate"]+"key.pem", l),
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
		fmt.Println("Error ListenAndServe")
		fmt.Println(err)
	}

}
