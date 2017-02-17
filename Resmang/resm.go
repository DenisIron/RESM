package resm

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	clientRes = make(map[string]string)
)
//type clientRes struct{
//	clientRes map[string]string
//}

func Server(w http.ResponseWriter, r *http.Request) {
	var (
		resources = []string{"r1", "r2", "r3", "r4", "r5"} //slice, содержащий неиспользуемые ресурсы
		//clientRes = make(map[string]string)
		//dat    = data{}
	)
	//clientRes := make(map[string]string)

	request := strings.Split(r.URL.Path[1:], "/")
	switch request[0] {
	case "allocate":
		client := request[1] //Имя клиента
		//allocateHand(resources, &clientRes, client,  w)//   (resourDealloc, clientResor)
		count := 0
		for i := range resources {
			if clientRes[resources[i]] == "" {
				clientRes[resources[i]] = client
				break
			} else {count++ }
		}
		if count == len(resources){
			fmt.Fprintf(w, "Out of resource")
		}

		for i := range resources {
			fmt.Fprintf(w, clientRes[resources[i]])
			fmt.Fprintf(w, "\n")
		}

	case "deallocate":
		//deallocateHand(resources, w)
		resForDealloc := request[1] // Ресурс, который необходимо изъять у клиента
		count := 0
		for key:= range clientRes {
			if key == resForDealloc {
				delete(clientRes, key)
				break
			} else {count++ }
			//delete(clientRes, resForDealloc)
			//fmt.Fprintf(w,"'%v':'%v' ", key, value)
		}
		if count == len(resources){
			fmt.Fprintf(w, "Not allocated")
		}

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
		list(resources, clientRes, w)

	case "reset":
		reset(resources, clientRes, w)
		//for key, value := range clientRes {
		//	delete(clientRes, key)
		//	fmt.Fprintf(w,"'%v':'%v' ", key, value)
		//}

	default:
		badRequest(w)
	}
}

//func allocateHand(resources []string, clientRes *map[string]string, client string,  w http.ResponseWriter) {
//	fmt.Fprintf(w, "\n")
//}

//func deallocateHand(resour []string, w http.ResponseWriter) {
//	fmt.Fprintf(w, resour[1])
//}

func list(resources []string, clientRes map[string]string, w http.ResponseWriter) {
	fmt.Fprintf(w, "'allocated': ")
	for key, value := range clientRes {
		fmt.Fprintf(w,"'%v':'%v' ", key, value)
	}

	fmt.Fprintf(w, " 'deallocated': ")
	for i := range resources {
		if clientRes[resources[i]] == "" {
			fmt.Fprintf(w, "'%s' ", resources[i])
		}
	}
}

func reset(resources []string, clientRes map[string]string, w http.ResponseWriter) {
	for key := range clientRes {
		delete(clientRes, key)
	}
	fmt.Fprintf(w, "RESET")
	for i := range resources {
		fmt.Fprintf(w, clientRes[resources[i]])
		fmt.Fprintf(w, "\n")
	}
}

func badRequest(w http.ResponseWriter) {
	w.Write([]byte("Bad Request")) //Неверный запрос
}