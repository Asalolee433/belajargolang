package main

import "fmt"

func main() {
	arr := [6]int{10, 20, 30, 40, 50, 60}

	s1 := arr[:] //ini untuk mengambil seluruh element
	fmt.Println("Ini adalah S1:", s1)

	s2 := arr[:3] //ini untuk mengambil index 0-2 (maksimum 3 jadi cuma diambil sampai 2)
	fmt.Println("Ini adalah s2:", s2)

	s3 := arr[2:] //ini untuk mengambil index ke 2-akhir (index ke 2 termasuk)
	fmt.Println("Ini adalah s3:", s3)

	s4 := arr[1:4] //ini untuk mengambil index ke 0-3 (4 tidak termasuk karna maksimum 4)
	fmt.Println("Ini adalah s4:", s4) //secara nilai akan otomatis berhenti di index ke 3 tapi untuk capacity itu 								sampai ke 4
}
