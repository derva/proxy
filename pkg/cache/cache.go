package cache

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/derva/proxy/pkg/logger"
)

var CacheRegister = make(map[string]http.Response)
var CacheRegisterCounter = make(map[string]uint8)
var CacheTreshold uint8 = 2

func CheckLastModified(a, b string) (update bool) {
	a1, _ := time.Parse(time.RFC1123, a)
	b1, _ := time.Parse(time.RFC1123, b)

	return a1.Before(b1)
}

func GetWrapper(w http.ResponseWriter, r *http.Request) http.Response {
	res, err := http.Get(r.URL.String())
	if err != nil {
		fmt.Println("Error while fetchin data from proxy")
	}

	return *res
}

func CacheService(w http.ResponseWriter, req *http.Request, l logger.Logger) (http.Response, error) {
	val := req.Header.Get("Cache-Control")

	if len(val) == 0 {
		return http.Response{}, nil
	}

	if strings.Contains(val, "no-store") {
		l.Log("Caching is disabled, no-store directive/instructoin", true)
		return http.Response{}, nil
	}

	var data http.Response

	if IsInCache(req.URL.String()) {
		data, _ = ReadFromCache(req.URL.String())
		if strings.Contains(val, "no-cache") {
			originRes, err := http.Head(req.URL.String())
			if err != nil {
				l.Log("Error while fetching HEAD data from proxy", true)
				return http.Response{}, nil
			}
			modified := CheckLastModified(originRes.Header.Get("Last-Modified"), data.Header.Get("Last-Modified"))

			if !modified {
				fmt.Println("Everything up to date :) ")
				return data, nil
			}

		}
	} else {
		data = GetWrapper(w, req)
	}

	if strings.Contains(val, "no-transform") {
		fmt.Println("No transform")
	}

	if ShouldStoreInCache(req.URL.String()) {
		StoreInCache(req.URL.String(), data)
	}

	return data, nil
}

func ShouldStoreInCache(k string) bool {
	CacheRegisterCounter[k]++
	return CacheRegisterCounter[k] > CacheTreshold
}

func StoreInCache(k string, v http.Response) {
	CacheRegister[k] = v
}

func PrintCache() {
	for k, v := range CacheRegister {
		fmt.Println("key value")
		fmt.Println(k)
		fmt.Println(v)
	}
}

func IsInCache(v string) bool {
	_, ok := CacheRegister[v]
	return ok
}

func ReadFromCache(v string) (http.Response, error) {
	val, ok := CacheRegister[v]

	if !ok {
		return http.Response{}, fmt.Errorf("Error: Empty response")
	}

	return val, nil
}
