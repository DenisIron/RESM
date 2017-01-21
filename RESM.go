package main

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	w.Write([]byte("Hello Wdfdf!!!"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}
