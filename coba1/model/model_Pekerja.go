package model

import (
	"coba1/node"
)

var DaftarPekerja [10]node.Pekerja
var Kuota int = 10
var Index int = 0

func CreatePekerja(emp node.Pekerja) bool {
	if Index < Kuota {
		DaftarPekerja[Index] = emp
		Index = Index + 1
		return true 
	}
	return false
}

func ReadPekerja() []node.Pekerja {
	return DaftarPekerja[0:Index]
}

func UpdatePekerja(emp node.Pekerja, id int) bool {
	for i := 0; i < Index; i++ {
		if DaftarPekerja[i].ID == id {
			DaftarPekerja[i] = emp
			return true 
		}
	}
	return false
}
 
func DeletePekerja(id int) bool {
	for i := 0; i < Index; i++ {
		if DaftarPekerja[i].ID == id {
			DaftarPekerja[i] = DaftarPekerja[Index-1]
			Index = Index - 1
			return true 
		}
	}
	return false
}