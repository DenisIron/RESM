package main

import (
	"fmt"
	"net/http"
	"strings"
)

//type data struct { //Структура для хранения данных клинет-ресурс
//	res    string
//	client string
//}

var (
	clientRes = make(map[string]string)
)

func main() {
	http.HandleFunc("/", Server)
	http.ListenAndServe(":9000", nil)
}

func Server(w http.ResponseWriter, r *http.Request) {
	var (
		resources = []string{"r1", "r2", "r3", "r4", "r5"} //slice, содержащий неиспользуемые ресурсы
		//dat    = data{}
	)
	//clientRes := make(map[string]string)

	request := strings.Split(r.URL.Path[1:], "/")
	switch request[0] {
	case "allocate":
		client := request[1] //Имя клиента
		//allocateHand(resources, w)//   (resourDealloc, clientResor)
		for i := range resources {
			if clientRes[resources[i]] == "" {
				clientRes[resources[i]] = client
				break
			}
		}
		for i := range resources {
			fmt.Fprintf(w, clientRes[resources[i]])
			fmt.Fprintf(w, "\n")
		}

	case "deallocate":
		//deallocateHand(resources, w) //

		resForDealloc := request[1] // Ресурс, который необходимо изъять у клиента
		for range resources {
			if clientRes[resForDealloc] != "" {
				clientRes[resForDealloc] = ""
				break
			}
		}
		for i := range resources {
			fmt.Fprintf(w, clientRes[resources[i]])
			fmt.Fprintf(w, "\n")
		}

	case "list":
		fmt.Fprintf(w, "'allocated'")
		for key, value := range clientRes {
			fmt.Fprintf(w, "Key:", key, "Value:", value)
		}
		fmt.Fprintf(w, " 'deallocated': ")
		for i := range resources {
			if clientRes[resources[i]] != "" {
				fmt.Fprintf(w, "'%s' ", resources[i])
			}
		}
	//list(resources, w) //  (clientResor, resourDealloc)

	case "reset":
		for i := range resources {
			clientRes[resources[i]] = ""
		}
	//reset(resources, w) //  (clientResor, chResourDealloc)

	default:
		badRequest(w)
	}
}

func allocateHand(resour []string, w http.ResponseWriter) {
	fmt.Fprintf(w, resour[0])
}

func deallocateHand(resour []string, w http.ResponseWriter) {
	fmt.Fprintf(w, resour[1])
}

func list(resour []string, w http.ResponseWriter) {
	fmt.Fprintf(w, resour[2])
}

func reset(resour []string, w http.ResponseWriter) {
	fmt.Fprintf(w, resour[3])
}

func badRequest(w http.ResponseWriter) {
	w.Write([]byte("Bad Request")) //Неверный запрос
}

/*
func resources(r *http.Request) {
	var (
		clientResor     = [][]string{}               //двумерный slice для сохранения данных клиент-ресурс
		resourDealloc   = []string{"r1", "r2", "r3"} //одномерный slice для сохранения данных по deallocated
		chResourDealloc = make(chan string)          // канал для данных
	)
}*/

func allocateHandle(w http.ResponseWriter, r *http.Request) {
	//если localhost:9000/allocate/alice то есть Path(2: по /), то берется 1 ресурс клиенту
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
	//dat := data{}
	client := strings.Split(r.URL.Path[1:], "/")[1] //Имя клиента
	//dat[0].client = client //Добавляем в базу данные о клиенте
	fmt.Fprintf(w, client) // dat[0].client)
	//resour[0] = ""

	//clientResor := append(clientResor[][], client)
	//clientResor[0][1] = resour[0]
	//fmt.Fprintf(w, clientResor[0][0])
	//fmt.Fprintf(w, clientResor[0][1])
	//resour := append(resour, "rN")

	//for i := range resour {
	//	if resour[i] != "" {

	//		resour[i] = ""
	//		resour = resour[i:]
	//		break
	//	}
	//}
	//for i := range resour {
	//	fmt.Fprintf(w, "%s\n", string(resour[i]))
	//}

	//Обработка отсутствия ресурсов:
	//для тестирования: resour = []string{"", "", ""}
	//count := 0
	//for i := range resour {
	//	if resour[i] != "" {
	//		break
	//	}
	//	count++
	//}
	//if count == len(resour) {
	//	fmt.Fprintf(w, "Out of resource")
	//}
}

//func deallocate(d http.ResponseWriter, r *http.Request) {
//	//clientResor = [][]string{}
//	//localhost:8080/deallocate/r1 то в канал добавляется r1
//
//	allPath := r.URL.Path[1:]
//	dealResource := strings.Split(allPath, "/")[1] //Ресурс, который нужно deallocate
//	fmt.Fprintf(d, dealResource)
//
//	for i := range resour {
//		fmt.Fprintf(d, string(resour[i]))
//	}
//	resour := append(resour, dealResource) //Добавляем в неипользуемые ресурсы ресурс
//	for i := range resour {
//		fmt.Fprintf(d, "%s\n", string(resour[i]))
//	}
//	//n := len(resour)
//	//fmt.Fprintf(d, string(resour[i]))
//	dat := data{}
//	dat[0].res = string(resour[0]) //Добавляем в базу данные об используемом ресурсе
//	dat[0].client = "i"            //Добавляем в базу данные о клиенте
//
//	count := 0
//	for i := 0; i < len(dat); i++ {
//		if dat[i].res != dealResource {
//			count++
//		}
////	}
//	if count == len(dat) {
//		fmt.Fprintf(d, "Not allocate")
//	}
//}
//
//func list(clientResor [][]string, resourDealloc []string) // может пригодится
//func list(l http.ResponseWriter, r *http.Request) {
//	// "allocated":{"r1":"alice"}, "deallocated":["r2","r3"]
//	l.Write([]byte("Resources Manager - RESM\n"))
//
//	str := data{{"", ""}}
//	str[0].res = string(resour[0])
//	str[0].client = "client"
//
//	fmt.Fprintf(l, "'allocated': ")
//	fmt.Fprintf(l, "{'%s':'%s'} ", str[0].res, str[0].client) //allocated":{"r1":"alice"}
//	/*for i := range dat {
//		if string(dat[i].res) != "" {
//			fmt.Fprintf(l, "{'%s':'%s'} ", dat[0].res, dat[0].client) //allocated":{"r1":"alice"}
//		}
//	}*/
//	fmt.Fprintf(l, "'deallocated':[ ")
//	for i := range resour {
//		if string(resour[i]) != "" {
//			fmt.Fprintf(l, "'%s' ", string(resour[i]))
//		}
//	}
//	fmt.Fprintf(l, "]")
//}
//
//func reset(res http.ResponseWriter, r *http.Request) {
//	//Все ресурсы из слайса в канал передать и удалить всех клиентов
//	/*	for resour := range clientResor {
//		chResourDealloc <- clientResor
//		clientResor := ""
//	}*/
//	// для тестирования: resour := []string{"r1", "", "r3"}
//
//	for i := range resour {
//		if string(resour[i]) == "" {
//			resour[i] = "r" //+ string(i)
//		}
//	}
//	for i := range resour {
//		fmt.Fprintf(res, "'%s' ", string(resour[i]))
//	}
//	str := data{}
//
//	for i := range str {
//		str[i].res = ""
//		str[i].client = ""
//	}
//}
