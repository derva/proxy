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
        fmt.Println("Error while fetching data from proxy")
    }

    return *res
}

func Cache(req *http.Request, header string, l logger.Logger) (http.Response, error) {
    var data http.Response
    l.Log("Cache function", true)

    data, _ = ReadFromCache(req.URL.String())

    if strings.Contains(header, "no-cache") {
        originRes, err := http.Head(req.URL.String())
        if err != nil {
            l.Log("Error while fetching HEAD data from proxy", true)
            return http.Response{}, nil
        }

        modified := CheckLastModified(originRes.Header.Get("Last-Modified"), data.Header.Get("Last-Modified"))

        if modified {
            newData, err := http.Get(req.URL.String())
            if err != nil {
                fmt.Println("Error fetcing new data")
            }
            data = *newData
            fmt.Println("Everything up to date, using cached version :) ")
        } else {
            fmt.Println("Everything up to date, using cached version :) ")
        }
    }

    return data, nil
}

func CacheService(w http.ResponseWriter, req *http.Request, l logger.Logger) (http.Response, error) {
    var data http.Response
    var err error

    val := strings.ToLower(req.Header.Get("Cache-Control"))

    if len(val) == 0 {
        fmt.Println("Cache-control isn't specified")
        return http.Response{}, nil
    }

    if strings.Contains(val, "no-store") {
        l.Log("Caching is disabled, no-store directive/instructoin", true)
        return http.Response{}, nil
    }

    if IsInCache(req.URL.String()) {
        data, err = Cache(req, val, l)
        if err != nil {
            l.Log("Error while getting data from cache", true)
        }
        return data, nil
    } else {
        data = GetWrapper(w, req)
    }

    if ShouldStoreInCache(req.URL.String()) {
        StoreInCache(req.URL.String(), data)
        l.Log(req.URL.String() + " has been successfully stored in cache", true);
    }

    return data, nil
}

func ShouldStoreInCache(k string) bool {
    CacheRegisterCounter[k]++
    return CacheRegisterCounter[k] > CacheTreshold
}

//maybe to check here if storing fail
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
