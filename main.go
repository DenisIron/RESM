package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Resources Manager - RESM"))
	resources(r)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)

}

func resources(r *http.Request) {
	var (
		clientResor     = [][]string{}      //двумерный slice для сохранения данных клиент-ресурс
		resourDealloc   = []string{}        //одномерный slice для сохранения данных по deallocated
		chResourDealloc = make(chan string) // канал для данных
	)
	path := r.URL.Path() // додумать как работать с путем
	//path := r.URL.Path(2)
	switch path {
	case "allocate":
		allocate(resourDealloc)
	case "deallocate":
		deallocate(clientResor, resourDealloc)
	case "list":
		list(clientResor, resourDealloc)
	case "reset":
		reset(clientResor, chResourDealloc)
	default:
		BadRequest()
	}
}

func allocate(resourDealloc []string) {
	//если localhost:80/allocate/alice то есть Path(2: по /), то берется 1 ресурс клиенту
	if len(resourDealloc) != 0 {
		//прописать логику
	} else {
		fmt.Println("Out of resources")
	}
}

func deallocate(clientResor [][]string, resourDealloc []string) {
	for range clientResor {
		//прописать логику
	}
	clientResor = [][]string{}
	//localhost:8080/deallocate/r1 то в канал добавляется r1

}

func list(clientResor [][]string, resourDealloc []string) {
	// "allocated":{"r1":"alice"}, "deallocated":["r2","r3"]
	fmt.Printf("allocaetd: {%s}, deallocated: %d\n", clientResor, resourDealloc)
}

func reset(clientResor [][]string, chResourDealloc chan string) {
	//Все ресурсы из слайса в канал передать и удалить всех клиентов
	for resour := range clientResor {
		chResourDealloc <- clientResor
		clientResor := ""
	}
}

func BadRequest() {
	//Ошибка
	fmt.Println("Bad Request")
}
