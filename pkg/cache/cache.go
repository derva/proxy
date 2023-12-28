package cache

import "fmt"

var CacheRegister = make(map[string]string)

func StoreInCache(k, v string) error {
	CacheRegister[k] = v
	return nil
}

func ReadFromCache(v string) bool {
	val, ok := CacheRegister[v]
	if ok {
		fmt.Println("value is ")
		fmt.Println(val)
		fmt.Println(CacheRegister[v])
		return true
	} else {
		return false
	}
}

func FetchFromCache(v string) {

}
