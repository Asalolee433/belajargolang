package main

import "fmt"

func main() {

	var nilai int
	fmt.Println("Masukkan Nilai anda")
	fmt.Scanln(&nilai)

	if nilai >= 90 {
		fmt.Println("Nilai anda A (Sangat Baik)")
	} else if nilai >= 75 {
		fmt.Println("Nilai anda B (Baik)")
	} else if nilai >= 60 {
		fmt.Println("Nilai anda C (Cukup)")
	} else {
		fmt.Println("Nilai anda D (Perlu perbaikan)")
	}

	//IF SHORT STATEMENT
	var angka int
	fmt.Println("Masukkan angka anda")
	fmt.Scanln(&angka)
	
	if v := angka*2; v> 10{
		fmt.Println("Hasil lebih dari 10", v)
	}
}