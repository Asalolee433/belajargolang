package main

import "fmt"

func main() {
	hari := "Sabtu"

	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari libur")
	}
}