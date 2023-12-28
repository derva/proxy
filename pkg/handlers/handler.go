package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/derva/proxy/pkg/logger"
)

func CheckLastModified(a, b string) (update bool) {
	a1, _ := time.Parse(time.RFC1123, a)
	b1, _ := time.Parse(time.RFC1123, b)

	return a1.Before(b1)
}

func HandleWrapper(w http.ResponseWriter, r *http.Request) {
	l := logger.LoadLogger("proxy.log", "/var/log/")

	encoding := Encoding(w, r, l)
	fmt.Println(encoding)

}
