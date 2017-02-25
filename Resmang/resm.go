package resm

import (
	"fmt"
	"net/http"
	"strings"
)

type ClientRes struct{
	Resources []string
	ClientResMap map[string]string
}

func Server(w http.ResponseWriter, r *http.Request) {
	struc := ClientRes{}
	struc.Resources = []string{"r1", "r2", "r3", "r4", "r5"} //slice, содержащий все ресурсы
	struc.ClientResMap = map[string]string{}

	//clientRes.fillingSliceOfRes()
	request := strings.Split(r.URL.Path[1:], "/")
	switch request[0] {
	case "allocate":
		client := request[1]
		statusBool := struc.Allocate(client)
		if !statusBool {
			fmt.Fprintf(w, "Out of resource")
		}

	case "deallocate":
		resForDealloc := request[1] // Ресурс, который необходимо изъять у клиента
		statusBool := struc.Deallocate(resForDealloc)
		if !statusBool {
			fmt.Fprintf(w, "Not allocated")
		}

	case "list":
		//	list()
		fmt.Fprintf(w, "'allocated': ")
		for key, value := range struc.ClientResMap {
			fmt.Fprintf(w,"'%v':'%v' ", key, value)
		}

		fmt.Fprintf(w, " 'deallocated': ")
		for i := range struc.Resources {
			if struc.ClientResMap[struc.Resources[i]] == "" {
				fmt.Fprintf(w, "'%s' ", struc.Resources[i])
			}
		}

	case "reset":
		struc.Reset()

	default:
		w.Write([]byte("Bad Request"))
	}
}

func (Struc *ClientRes) Allocate(client string) bool {
	for i := range Struc.Resources {
		if Struc.ClientResMap[Struc.Resources[i]] == "" {
			Struc.ClientResMap[Struc.Resources[i]] = client
			return true
		}
	}
	return false
}

func (Struc *ClientRes) Deallocate(resForDealloc string) bool {
	if Struc.ClientResMap[resForDealloc] != "" {
		delete(Struc.ClientResMap, resForDealloc)
		return true
	} else {
		return false
	}
}

func (Struc *ClientRes) Reset() {
	for keyResources := range Struc.ClientResMap {
		delete(Struc.ClientResMap, keyResources)
	}
}
/*
func list() {
	fmt.Fprintf(w, "'allocated': ")
	for key, value := range clientRes {
		//fmt.Fprintf(w,"'%v':'%v' ", key, value)
	}

	fmt.Fprintf(w, " 'deallocated': ")
	for i := range resources {
		if clientRes[resources[i]] == "" {
			fmt.Fprintf(w, "'%s' ", resources[i])
		}
	}
}
*/
//func (res *clientRes) fillingSliceOfRes(){
//	res.resources = []string{"r1", "r2", "r3", "r4", "r5"}
//}

//func allocate(resources []string, client string,  w http.ResponseWriter) {
//	fmt.Fprintf(w, "\n")
//}
/*
func (clientRes *clientRes) allocate(client string,  w http.ResponseWriter) {
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

//func deallocate(resour []string, w http.ResponseWriter) {

//}