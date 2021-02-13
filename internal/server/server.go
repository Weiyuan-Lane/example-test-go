package server

import (
	"fmt"
	"net/http"

	"github.com/weiyuan-lane/example-test-go/internal/server/utils/queryparams"
)

// Adapted from https://gobyexample.com/http-servers

func Run() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/test-page-params", testPageParams)

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func testPageParams(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	queryParams := queryparams.New(params)

	page, _ := queryParams.Page()
	perPage, _ := queryParams.PerPage()

	fmt.Fprintf(w, "page: %v\n", page)
	fmt.Fprintf(w, "perPage: %v\n", perPage)
}
