package main

import (
	"fmt"
	"net/http"
)

var (
	port = ":80"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/home", homeHandler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
