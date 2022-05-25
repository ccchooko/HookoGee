package main

import (
	"fmt"
	"github.com/ccchooko/HookoGee/gee"
	"log"
	"net/http"
)
func main() {
	r := gee.New()
	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)

	log.Fatal(r.Run(":9999"))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Url.Path = %q\n", req.URL.Path)
}
