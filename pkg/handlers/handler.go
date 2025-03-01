package handlers

import (
	"fmt"
	"net/http"

	"github.com/derva/proxy/pkg/cache"
	"github.com/derva/proxy/pkg/logger"
)

func GetResponse(w http.ResponseWriter, r *http.Request) *http.Response {
	res, err := http.Get(r.URL.String())
	if err != nil {
		fmt.Println("Error while fetching data from proxy")
	}
	//defer res.Body.Close()

	//body, _ := io.ReadAll(res.Body)

	return res
}

func HandleWrapper(w http.ResponseWriter, req *http.Request) {
	l := logger.LoadLogger("proxy.log", "/var/log/")

    test, err := cache.CacheService(w, req, l)
    if err != nil {
        l.Log("nema errora", true);
    }

    l.Log("before status testing", true)
    l.Log(test.Status, true);

	Encoding(w, req, l)

}
