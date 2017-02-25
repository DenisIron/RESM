package main

import (
	"net/http"
	"Resmang"
)

func main() {
	/*
	"github.com/julienschmidt/httprouter"
	router := httprouter.New()
	router.GET("/", indexHandler)
	http.ListenAndServe(":8080", router)
	*/

	http.HandleFunc("/", resm.Server)
	http.ListenAndServe(":9000", nil)
}
