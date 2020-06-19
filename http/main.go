package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/test2", HelloServer)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello2, %s!", r.URL.Path[1:])
		fmt.Printf("Hello2, %s!", r.URL.Path[1:])
	})
	fmt.Printf("About to start:")
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET, %s!", r.URL.Path[1:])
	case "POST":
		// curl -X POST http://localhost:8080/test2 -d "fgds  sdfgsdgf dsfgsdfg  sdfgsdfg"
		b, _ := ioutil.ReadAll(r.Body)
		s := string(b)
		fmt.Fprintf(w, "POST, %s! %v", r.URL.Path[1:], s)
	}
	//fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])

	//Homework
	// go get github.com/gin-gonic/gin
}
