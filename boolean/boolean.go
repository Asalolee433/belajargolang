package main

import("fmt")

func main() {
	//Deklarasi variabel boolean 
	var isRaining bool = true
	var isSunny bool = false
	isLoggedIn := true
	hasPermission := false

	//Menampilkan nilai boolean
	fmt.Println("Apakah hujan hari ini?:", isRaining)
	fmt.Println("Apakah terang hari ini?:", isSunny)
	fmt.Println("Apakah udah login?:", isLoggedIn)
	fmt.Println("Apakah punya akses?:", hasPermission)
}
