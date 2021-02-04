package server

import (
	"os"
)

var (
	port = os.Getenv("PORT")
)

func Run() {
	// http.HandleFunc("/hello", hello)
	// http.HandleFunc("/headers", headers)

	// http.ListenAndServe(port, nil)
}
