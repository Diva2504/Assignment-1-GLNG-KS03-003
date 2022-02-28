package main

import (
	"fmt"
	"os"
)

type kelas struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (k kelas) tampil() {
	fmt.Printf(" Nama : %s \n Alamat : %s \n Pekerjaan : %s \n Alasan memilih kelas Golang : %s \n", k.Nama, k.Alamat, k.Pekerjaan, k.Alasan)
}

func main() {
	argument := os.Args
	faris := kelas{Nama: "1", Alamat: "1", Pekerjaan: "1", Alasan: "1"}
	rakha := kelas{Nama: "2", Alamat: "2", Pekerjaan: "2", Alasan: "2"}
	diva := kelas{Nama: "3", Alamat: "3", Pekerjaan: "3", Alasan: "3"}
	daffa := kelas{Nama: "4", Alamat: "4", Pekerjaan: "4", Alasan: "4"}
	yoga := kelas{Nama: "5", Alamat: "5", Pekerjaan: "5", Alasan: "5"}
	azhar := kelas{Nama: "6", Alamat: "6", Pekerjaan: "6", Alasan: "6"}
	haiqal := kelas{Nama: "7", Alamat: "7", Pekerjaan: "7", Alasan: "7"}
	luqman := kelas{Nama: "8", Alamat: "8", Pekerjaan: "8", Alasan: "8"}
	ruben := kelas{Nama: "9", Alamat: "9", Pekerjaan: "9", Alasan: "9"}
	denada := kelas{Nama: "10", Alamat: "10", Pekerjaan: "10", Alasan: "10"}

	if len(argument) == 2 {
		switch argument[1] {
		case "1":
			faris.tampil()
		case "2":
			rakha.tampil()
		case "3":
			diva.tampil()
		case "4":
			daffa.tampil()
		case "5":
			yoga.tampil()
		case "6":
			azhar.tampil()
		case "7":
			haiqal.tampil()
		case "8":
			luqman.tampil()
		case "9":
			ruben.tampil()
		case "10":
			denada.tampil()
		}
	} else {
		fmt.Println("too much argument")
	}

}
