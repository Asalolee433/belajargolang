package main

import "fmt"

func main() {
	var angka1, angka2 int
	//Input dari pengguna
	fmt.Println("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)
	fmt.Println("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)
	//Latihan Perbandingan
	fmt.Println("\n=== Hasil Perbandingan ===")
	fmt.Printf("%d == %d ? %v\n", angka1, angka2, angka1 == angka2)
	fmt.Printf("%d != %d ? %v\n", angka1, angka2, angka1 != angka2)
	fmt.Printf("%d > %d ? %v\n", angka1, angka2, angka1 > angka2)
	fmt.Printf("%d < %d ? %v\n", angka1, angka2, angka1 < angka2)
	fmt.Printf("%d >= %d ? %v\n", angka1, angka2, angka1 >= angka2)
	fmt.Printf("%d <= %d ? %v\n", angka1, angka2, angka1 <= angka2)
}