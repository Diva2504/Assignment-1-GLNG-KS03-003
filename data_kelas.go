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
	faris := kelas{Nama: "Ahmad Faris", Alamat: "1", Pekerjaan: "1", Alasan: "1"}
	rakha := kelas{Nama: "Rakha Rizqllah Pratama Saputra", Alamat: "2", Pekerjaan: "2", Alasan: "2"}
	diva := kelas{Nama: "I Gede Diva Dwijayana", Alamat: "Karangasen, Bali", Pekerjaan: "Mahasiswa", Alasan: "karena penasaran dengan bahasa pemrograman Golang, ingin mempelajarinya, dan membuat sebuah aplikasi dari Golang"}
	daffa := kelas{Nama: "Muhammad Daffa Haryadi Putra", Alamat: "4", Pekerjaan: "4", Alasan: "4"}
	yoga := kelas{Nama: "M yoga irvandi", Alamat: "Sumatra Selatan", Pekerjaan: "Mahasiswa", Alasan: "karena Golang merupakan bahasa pemrograman masih tergolong baru,yg menarik untuk di pelajari, bagi saya."}
	azhar := kelas{Nama: "Wafianda Azhar", Alamat: "6", Pekerjaan: "6", Alasan: "6"}
	haiqal := kelas{Nama: "Muhamad Haiqal Malik", Alamat: "7", Pekerjaan: "7", Alasan: "7"}
	luqman := kelas{Nama: "Luqman Fauzi", Alamat: "8", Pekerjaan: "8", Alasan: "8"}
	ruben := kelas{Nama: "Rubenhard Manat Hasiholan Lumbantobing", Alamat: "9", Pekerjaan: "9", Alasan: "9"}
	denada := kelas{Nama: "Denada Ramschie", Alamat: "10", Pekerjaan: "10", Alasan: "10"}

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
		default:
			fmt.Println("No data....")
		}
	} else {
		fmt.Println("Invalid Argument / Empty argument")
	}

}
