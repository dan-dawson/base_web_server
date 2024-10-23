package main

import "net/http"

var (
	port = ":80"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	http.HandleFunc("/home", homeHandler)
	http.ListenAndServe(port, nil)
}
