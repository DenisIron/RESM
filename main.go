package main

import (
	"fmt"
	"net/http"
	"strings"
)

var st []int
var (
	resour      = []string{"r1", "r2", "r3"} //слайс содержащий неиспользуемые ресурсы
	clientResor = [][]string{}               //слайс слайсов содержащий данные клиент-ресурс
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
}*/

func allocate(w http.ResponseWriter, r *http.Request) {
	//если localhost:80/allocate/alice то есть Path(2: по /), то берется 1 ресурс клиенту
	/*	if len(resourDealloc) != 0 {
			res := resourDealloc[0]
			resourDealloc[0] = ""
			clientResor = res
		} else {
			fmt.Println("Out of resources")
		}*/
	//s := append(st, 3)
	w.Write([]byte("Resources Manager - RESM\n"))
	path := r.URL.Path[1:]
	//fmt.Fprintf(w, "Команда: %s!\n", path)

	/*for i := range dd {
		fmt.Fprintf(w, string(dd[i]))
		if string(dd[i]) == "/" {
			client := strings.Split(string(dd), "/")
			fmt.Fprintf(w, client[1])
			break
		}
	}*/
	allPath := []byte(path)
	client := strings.Split(string(allPath), "/")
	fmt.Fprintf(w, client[1])
	//resour := append(resour, "rN")

	for i := range resour {
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, string(resour[i]))
	}

}

func deallocate(d http.ResponseWriter, r *http.Request) {
	/*for range clientResor {
		//прописать логику
	}
	clientResor = [][]string{}*/
	//localhost:8080/deallocate/r1 то в канал добавляется r1
	d.Write([]byte("Resources Manager - RESM\n"))
	path := r.URL.Path[1:]
	fmt.Fprintf(d, "Команда: %s!\n", path)

	for i := range resour {
		if resour[i] != "" {
			resour[i] = ""
			resour = resour[i:]
			break
		}
	}
	for i := range resour {
		fmt.Fprintf(d, "\n")
		fmt.Fprintf(d, string(resour[i]))
	}
	//n := len(resour)
	//fmt.Fprintf(d, string(resour[i]))

}

//func list(clientResor [][]string, resourDealloc []string) // может пригодится
func list(l http.ResponseWriter, r *http.Request) {
	// "allocated":{"r1":"alice"}, "deallocated":["r2","r3"]
	//	fmt.Fprintf("allocaetd: {%s}, deallocated: %d\n", clientResor, resourDealloc)*/
	l.Write([]byte("Resources Manager - RESM\n"))
	path := r.URL.Path[1:]
	fmt.Fprintf(l, "Команда: %s!\n", path)

	for i := range resour {
		fmt.Fprintf(l, "\n")
		fmt.Fprintf(l, string(resour[i]))
	}
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
func head() {
	w.Write([]byte("Resources Manager - RESM\n"))
	path := r.URL.Path[1:]
	fmt.Fprintf(w, "Команда: %s!\n", path)
}*/

/*
func BadRequest() {
	//Ошибка
	fmt.Println("Bad Request")
}*/
