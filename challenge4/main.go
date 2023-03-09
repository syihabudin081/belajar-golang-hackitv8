package main

import (
	"fmt"
	"os"
)

func main() {
	

	type biodata struct {
		absen     int
		nama      string
		alamat    string
		pekerjaan string
		alasan    string
	}

	var biodataTeman = []biodata{
		{1, "Agus", "Jln Merdeka", "Programmer", "Ingin Belajar Golang"}, {2, "Asep", "Jln Merdeka 2", "Junior Programmer", "Ingin Belajar Gin"}, {3, "Joko", "Jln Merdeka 3", "Senior Programmer", "Ingin Belajar Golang"},
	}

	args := os.Args



	absen := args[1]

	for _,teman := range biodataTeman {
		if fmt.Sprint(teman.absen) == absen {
			fmt.Println("Data teman dengan absen", absen, ":")
			fmt.Println("Nama    :", teman.nama)
			fmt.Println("Alamat  :", teman.alamat)
			fmt.Println("Pekerjaan:", teman.pekerjaan)
			fmt.Println("Alasan  :", teman.alasan)
			return
		}
	}

	fmt.Println("Data Dengan Absen : ", absen, " Tidak Ditemukan")
}