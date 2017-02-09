package main

import (
	"fmt"
	"net/http"
)

var st []int
var (
	resour = []string{"r1", "r2", "r3"}
)

func main() {
	http.HandleFunc("/allocate/", allocate)
	http.HandleFunc("/deallocate/", deallocate)
	http.HandleFunc("/list/", list)
	http.HandleFunc("/reset/", reset)
	http.ListenAndServe(":9000", nil)

}

/*
func resources(r *http.Request) {
	var (
		clientResor     = [][]string{}               //двумерный slice для сохранения данных клиент-ресурс
		resourDealloc   = []string{"r1", "r2", "r3"} //одномерный slice для сохранения данных по deallocated
		chResourDealloc = make(chan string)          // канал для данных
	)
	path := r.URL.Path() // додумать как работать с путем
	//path := r.URL.Path(2)
	switch path {
	case "allocate":
		allocate(resourDealloc, clientResor)
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
*/
func allocate(w http.ResponseWriter, r *http.Request) {
	//если localhost:80/allocate/alice то есть Path(2: по /), то берется 1 ресурс клиенту
	/*	if len(resourDealloc) != 0 {
			res := resourDealloc[0]
			resourDealloc[0] = ""
			clientResor = res
			//прописать логику
		} else {
			fmt.Println("Out of resources")
		}*/
	//s := append(st, 3)
	w.Write([]byte("Resources Manager - RESM\n"))
	path := r.URL.Path[1:]
	fmt.Fprintf(w, "Команда: %s!\n", path)
	fmt.Fprintf(w, string(resour[0]))
	s := append(resour, "2")
	fmt.Fprintf(w, string(s[3]))
	resour[0] = ""
	fmt.Fprintf(w, string(resour[0]))

}

func deallocate(d http.ResponseWriter, r *http.Request) {
	/*for range clientResor {
		//прописать логику
	}
	clientResor = [][]string{}*/
	//localhost:8080/deallocate/r1 то в канал добавляется r1
	w.Write([]byte("Resources Manager - RESM\n"))
	path := r.URL.Path[1:]
	fmt.Fprintf(w, "Команда: %s!\n", path)

}

//func list(clientResor [][]string, resourDealloc []string) // может пригодится
func list(l http.ResponseWriter, r *http.Request) {
	// "allocated":{"r1":"alice"}, "deallocated":["r2","r3"]
	//	fmt.Fprintf("allocaetd: {%s}, deallocated: %d\n", clientResor, resourDealloc)*/
	fmt.Fprintf(l, "list")
}

func reset(res http.ResponseWriter, r *http.Request) {
	//Все ресурсы из слайса в канал передать и удалить всех клиентов
	/*	for resour := range clientResor {
		chResourDealloc <- clientResor
		clientResor := ""
	}*/
	fmt.Fprintf(res, "reset")
}

/*
func BadRequest() {
	//Ошибка
	fmt.Println("Bad Request")
}*/
