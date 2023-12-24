package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Handle(w http.ResponseWriter, r *http.Request) { 
	http.Redirect(w, r, r.URL.Host, http.StatusUseProxy)
}

func main() {
	fmt.Println("Server is starting... ")

	router := mux.NewRouter();
	
	router.HandleFunc("/", Handle)

	s := http.Server{
		Addr: ":8080",
		Handler: router,
	}

	err := s.ListenAndServe();
	if err != nil {
		fmt.Println("Error");
	}

}