package server

import (
	"fmt"
	"net/http"

	"github.com/Weiyuan-Lane/example-test-go/internal/server/utils/queryparams"
)

// Adapted from https://gobyexample.com/http-servers

func Run() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/page-params", pageParams)

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

func pageParams(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	queryParams := queryparams.New(params)
	RenderPageParamsResponse(w, queryParams)
}

func RenderPageParamsResponse(w http.ResponseWriter, queryParams queryparams.PageParams) {
	page, _ := queryParams.Page()
	perPage, _ := queryParams.PerPage()

	fmt.Fprintf(w, "page: %v\n", page)
	fmt.Fprintf(w, "perPage: %v\n", perPage)
}
