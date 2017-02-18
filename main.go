package main

import (
	"net/http"
	"Resmang"
)

func main() {
	http.HandleFunc("/", resm.Server)
	http.ListenAndServe(":9000", nil)
}
