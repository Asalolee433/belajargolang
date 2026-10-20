package main

import "fmt"

func main(){
	mahasiswa := map[string]string{
		"nama" : "felix",
		"kelas" : "A",
	}
	fmt.Println("Nama :", mahasiswa["nama"]) //Nama : felis
	fmt.Println("Kelas :", mahasiswa["kelas"]) //Kelas : A

	mahasiswa["jurusan"] = "Informatika"
	fmt.Println("Jurusan :", mahasiswa["jurusan"])
}
