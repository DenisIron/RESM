package resm

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	clientRes = make(map[string]string)
)

type clientRes2 struct{
	resources2 []string
 	clientResMap2 map[string]string
}

func Server(w http.ResponseWriter, r *http.Request) {
	var (
		resources = []string{"r1", "r2", "r3", "r4", "r5"} //slice, содержащий все ресурсы
	)
	//clientRes.fillingSliceOfRes()
	request := strings.Split(r.URL.Path[1:], "/")
	switch request[0] {
	case "allocate":
		client := request[1] //Имя клиента
		//allocateHand(resources, client,  w)//   (resourDealloc, clientResor)
/*
clientRes.allocate(client,  w)
*/
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

	case "deallocate":
		//deallocateHand(resources, w)
		resForDealloc := request[1] // Ресурс, который необходимо изъять у клиента
		count := 0

		for range resources {
			if clientRes[resForDealloc] != "" {
				delete(clientRes, clientRes[resForDealloc])
				break
			} else {
				count++
			}
		}
		//for key:= range clientRes {
		//	if key == resForDealloc {
		//		delete(clientRes, key)
		//		break
		//	} else {count++ }
		//}
		if count == len(resources){
			fmt.Fprintf(w, "Not allocated")
		}

		//for range resources {
		//	if clientRes[resForDealloc] != "" {
		//		delete(clientRes, clientRes[resForDealloc])
		//		break
		//	}
		//}

	case "list":
		list(resources, clientRes, w)

	case "reset":
		reset(resources, clientRes, w)

	default:
		badRequest(w)
	}
}

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
	//fmt.Fprintf(w, "RESET")
}

func badRequest(w http.ResponseWriter) {
	w.Write([]byte("Bad Request")) //Неверный запрос
}

//func (res *clientRes2) fillingSliceOfRes(){
//	res.resources2 = []string{"r1", "r2", "r3", "r4", "r5"}
//}

//func allocateHand(resources []string, client string,  w http.ResponseWriter) {
//	fmt.Fprintf(w, "\n")
//}
/*
func (clientRes *clientRes) allocateHand(client string,  w http.ResponseWriter) {
	count := 0
	for i := range clientRes.resources {
		if clientRes.clientRes[resources[i]] == "" {
			clientRes[resources[i]] = client
				break
			} else {count++ }
		}
		if count == len(resources){
			fmt.Fprintf(w, "Out of resource")
		}
	fmt.Fprintf(w, "\n")
}
*/

//func deallocateHand(resour []string, w http.ResponseWriter) {

//}