package main

import "fmt"

func main(){
	//Deklarasi string dengan kutip
	nama := "Antony"
	pesan := "Selamat Datang di aplikasi kami!"


	//Deklarasi string dengan backtick (raw string)
	paragraf := `Halo, ini adalah contoh multi-line 
	string di Golang`

	//Menampilkan nilai string
	fmt.Println("Nama:", nama)
	fmt.Println("Pesan:", pesan)
	fmt.Println("Ucapan:\n", paragraf)
}