package resm

import (
	//"fmt"
	//"net/http"
	//"strings"
	"testing"
)

func TestOfAllocateHand(t *testing.T) {
	struc := ClientRes{}
	struc.Resources = []string{"r1", "r2", "r3", "r4", "r5"}
	struc.ClientResMap = map[string]string{"r1": "", "r2": "alice"}
	struc.Allocate("bob")
	if struc.ClientResMap[struc.Resources[0]] != "bob" {
		t.Fatal("AllocateHandle does not work correctly - Resourse 1 must be bob")
	}
	struc.ClientResMap = map[string]string{"r1": "bob", "r2": "alice", "r3": "", "r4": ""}
	struc.Allocate("alice")
	struc.Allocate("bob")
	if (struc.ClientResMap[struc.Resources[2]] != "alice") && (struc.ClientResMap[struc.Resources[3]] != "bob") {
		if (struc.ClientResMap[struc.Resources[2]] != "alice") && (struc.ClientResMap[struc.Resources[3]] == "bob") {
			t.Fatal("AllocateHandle does not work correctly - Resourse 3 must be alice")
		}
		if (struc.ClientResMap[struc.Resources[2]] == "alice") && (struc.ClientResMap[struc.Resources[3]] != "bob") {
			t.Fatal("AllocateHandle does not work correctly - Resourse 4 must be bob")
		} else {
			t.Fatal("AllocateHandle does not work correctly - Resourse 4 must be bob")
		}
		t.Fatal("AllocateHandle does not work correctly - Resourse 1 must be bob")
	}

}

func TestOfDeallocateHand(t *testing.T) {
	struc := ClientRes{}
	struc.Resources = []string{"r1", "r2", "r3", "r4", "r5"}
	struc.ClientResMap = map[string]string{"r1": "", "r2": "alice"}

	struc.Deallocate("r2")
	if struc.ClientResMap[struc.Resources[1]]!=""{
		t.Fatal("DeallocateHandle does not work correctly - Resourse 2 must be nil")
	}
	struc.ClientResMap = map[string]string{"r1": "bob", "r2": "alice", "r3": "", "r4": ""}
	struc.Deallocate("r2")
	if struc.ClientResMap[struc.Resources[1]]!=""{
		t.Fatal("DeallocateHandle does not work correctly - Resourse 2 must be nil")
	}
	struc.Deallocate("r1")
	if struc.ClientResMap[struc.Resources[0]]!=""{
		t.Fatal("DeallocateHandle does not work correctly - Resourse 1 must be nil")
	}


}

func TestOfList(t *testing.T) {

}

func TestOfReset(t *testing.T) {
	struc := ClientRes{}
	struc.Resources = []string{"r1", "r2", "r3", "r4", "r5"}
	struc.ClientResMap = map[string]string{"r1": "bob", "r2": "alice", "r3": "", "r4": "bob", "r5": "jack",}
	struc.Reset()
	for struc.Resources{
		if struc.ClientResMap[struc.Resources] != ""{
			t.Fatal("ResetHandle does not work correctly - All resources must be nil")
			break
		}
	}


}
