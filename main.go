package main

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	resour = []string{"r1", "r2", "r3"} //slice, содержащий неиспользуемые ресурсы
	//clientResor = [][]string{}               //слайс слайсов содержащий данные клиент-ресурс
)

type data []struct { //Структура для хранения данных клинет-ресурс
	res, client string
}

func main() {
	http.HandleFunc("/allocate/", allocate)
	http.HandleFunc("/deallocate/", deallocate)
	http.HandleFunc("/list/", list)
	http.HandleFunc("/reset/", reset)
	http.HandleFunc("/", badRequest)
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
	/*for i := range dd {
		fmt.Fprintf(w, string(dd[i]))
		if string(dd[i]) == "/" {
			client := strings.Split(string(dd), "/")
			fmt.Fprintf(w, client[1])
			break
		}
	}*/

	allPath := []byte(r.URL.Path[1:])                //Весь path после слеша
	client := strings.Split(string(allPath), "/")[1] //Имя клиента
	dat := data{}
	dat[0].res = string(resour[0]) //Добавляем в базу данные об используемом ресурсе
	dat[0].client = client         //Добавляем в базу данные о клиенте
	fmt.Fprintf(w, dat[0].client)
	resour[0] = ""

	//clientResor := append(clientResor[][], client)
	//clientResor[0][1] = resour[0]
	//fmt.Fprintf(w, clientResor[0][0])
	//fmt.Fprintf(w, clientResor[0][1])
	//resour := append(resour, "rN")
	for i := range resour {
		if resour[i] != "" {
			resour[i] = ""
			resour = resour[i:]
			break
		}
	}
	for i := range resour {
		fmt.Fprintf(w, "%s\n", string(resour[i]))
	}

}

func deallocate(d http.ResponseWriter, r *http.Request) {
	/*for range clientResor {
		//прописать логику
	}
	clientResor = [][]string{}*/
	//localhost:8080/deallocate/r1 то в канал добавляется r1
	d.Write([]byte("Resources Manager - RESM\n"))
	allPath := r.URL.Path[1:]
	dealResource := strings.Split(string(allPath), "/")[1] //Ресурс, который нужно deallocate
	fmt.Fprintf(d, dealResource)

	for i := range resour {
		fmt.Fprintf(d, string(resour[i]))
	}
	resour := append(resour, dealResource) //Добавляем в неипользуемые ресурсы ресурс
	for i := range resour {
		fmt.Fprintf(d, "%s\n", string(resour[i]))
	}
	//n := len(resour)
	//fmt.Fprintf(d, string(resour[i]))

}

//func list(clientResor [][]string, resourDealloc []string) // может пригодится
func list(l http.ResponseWriter, r *http.Request) {
	// "allocated":{"r1":"alice"}, "deallocated":["r2","r3"]
	l.Write([]byte("Resources Manager - RESM\n"))

	str := data{{"", ""}}
	str[0].res = string(resour[0])
	str[0].client = "client"

	fmt.Fprintf(l, "'allocated': ")
	fmt.Fprintf(l, "{'%s':'%s'} ", str[0].res, str[0].client) //allocated":{"r1":"alice"}
	/*for i := range dat {
		if string(dat[i].res) != "" {
			fmt.Fprintf(l, "{'%s':'%s'} ", dat[0].res, dat[0].client) //allocated":{"r1":"alice"}
		}
	}*/
	fmt.Fprintf(l, "'deallocated':[ ")
	for i := range resour {
		if string(resour[i]) != "" {
			fmt.Fprintf(l, "'%s' ", string(resour[i]))
		}
	}
	fmt.Fprintf(l, "]")
}

func reset(res http.ResponseWriter, r *http.Request) {
	//Все ресурсы из слайса в канал передать и удалить всех клиентов
	/*	for resour := range clientResor {
		chResourDealloc <- clientResor
		clientResor := ""
	}*/
	// для тестирования: resour := []string{"r1", "", "r3"}

	for i := range resour {
		if string(resour[i]) == "" {
			resour[i] = "r" //+ string(i)
		}
	}
	for i := range resour {
		fmt.Fprintf(res, "'%s' ", string(resour[i]))
	}
	str := data{}

	for i := range str {
		str[i].res = ""
		str[i].client = ""
	}
}

/*
func head() {
	w.Write([]byte("Resources Manager - RESM\n"))
}*/

func badRequest(bad http.ResponseWriter, r *http.Request) {
	bad.Write([]byte("Bad Request")) //Неверный запрос
}
