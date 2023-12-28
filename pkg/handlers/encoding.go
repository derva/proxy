package handlers

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/derva/proxy/pkg/logger"
)

func PreferredEncodingAlgorithm() string {
	file, err := os.Open("conf.conf")
	if err != nil {
		fmt.Println("Error trying to open conf file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "PREFERRED_ENCODING_ALGORITHM") {
			value := strings.Split(line, "=")[1]
			return value
		}
	}
	return "nil"

}

func ChooseAlgorithm(algo string, l logger.Logger) string {
	prefered_algo := PreferredEncodingAlgorithm()
	if strings.Contains(algo, prefered_algo) {
		l.Log("Using preferred algorithm "+prefered_algo, true)
		return prefered_algo
	} else {
		algos := strings.Split(algo, " ")
		l.Log("Using algorithm "+algos[0], true)
		return algos[0]
	}
}

func Encoding(w http.ResponseWriter, r *http.Request, l logger.Logger) string {
	resp, err := http.Get(r.URL.String())
	if err != nil {
		l.Log("Error fetching the data from host", true)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	bodyString := string(body)

	val, ok := r.Header["Accept-Encoding"]

	if ok {
		l.Log("Encoding enabled ...", false)
	}

	algo := ChooseAlgorithm(val[0], l)

	switch algo {
	case "gzip":
		l.Log("GZIP encoding ...", true)
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		gz.Write([]byte(bodyString))
		break
	}
	return "Encoding finished using " + algo + " algorithm."
}
