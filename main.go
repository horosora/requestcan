package main

import (
	"log"
	"fmt"
	"strings"
	"time"
	"net/http"
	"net/http/httputil"
)

var (
	history []string
	historySize int = 5
)

func addHistory(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal(err)
	}

	if len(history) < historySize {
		history = append(history, string(dump))
	} else {
		history = append(history[1:], string(dump))
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(fmt.Sprintln(time.Now())))
}

func showHistory(res http.ResponseWriter, req *http.Request) {
	delimiter := "\n--------------------------------\n"
	data := delimiter + strings.Join(history, delimiter) + delimiter

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(data))
}

func clearHistory(res http.ResponseWriter, req *http.Request) {
	history = nil

	res.WriteHeader(http.StatusOK)
	res.Write([]byte("success\n"))
}

func main() {
	http.HandleFunc("/", addHistory)
	http.HandleFunc("/show", showHistory)
	http.HandleFunc("/clear", clearHistory)

	http.ListenAndServe(":8000", nil)
}
