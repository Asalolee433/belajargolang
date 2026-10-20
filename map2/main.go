package main

import "fmt"

func main() {
	mahasiswa := map[string]int{
		"Goebbels": 90,
		"Hitler":   80,
		"Wilhelm":  70,
	}
	fmt.Println(mahasiswa)//map[Goebbels:90 Hitler:80 Wilhelm:70]
}
