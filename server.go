package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Logger struct {
	LogFileName string
	Location    string
}

func (e *Logger) Log(s string) {
	f, err := os.OpenFile(e.Location+e.LogFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error Log() ")
		fmt.Println(err)
	}

	f.WriteString("[ " + time.Now().Format("01-02-2001 15:04:05") + " ] - " + s + "\n")
	fmt.Println(s)
}

func Handle(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, r.URL.Host, http.StatusUseProxy)
}

func loadCertificate(certFile, keyFile string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Certificates are loaded!")
	return cert
}

func main() {
	l := Logger{
		LogFileName: "proxy.log",
		Location:    "/var/log/",
	}

	l.Log("Server is starting ...")

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading dotenv files!")
	}

	CertificatePath := os.Getenv("CERTIFICATE_PATH")

	router := mux.NewRouter()

	tls := &tls.Config{
		Certificates: []tls.Certificate{
			loadCertificate(CertificatePath+"cert.pem", CertificatePath+"key.pem"),
		},
	}

	router.HandleFunc("/", Handle)

	s := http.Server{
		Addr:      ":8080",
		Handler:   router,
		TLSConfig: tls,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println("Error ListenAndServe")
	}

}
