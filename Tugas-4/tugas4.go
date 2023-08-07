package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Soal 1
	var batas = 20

	for i := 1; i <= batas; i++ {
		var I = strconv.Itoa(i)

		if i%2 == 1 {
			if i%3 == 0 {
				fmt.Println(I, " - I Love Coding")
			} else {
				fmt.Println(I, " - Santai")
			}
		} else {
			fmt.Println(I, " - Berkualitas")
		}
	}
	//Soal 2
	var tinggi = 7
	var pagar = ""
	for z := 1; z <= tinggi; z++ {
		pagar += "#"
		fmt.Println(pagar)
	}
	//Soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	var kalimatJadi = kalimat[2] + " " + kalimat[3] + " " + kalimat[4] + " " + kalimat[5] + " " + kalimat[6]
	fmt.Println(kalimatJadi)

	//Soal 4
	var index = 1
	var sayuran = []string{"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun"}

	for z := 0; z < len(sayuran); z++ {
		fmt.Println(index, sayuran[z])
		index++
	}

	//Soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}
	for a := 0; a < 1; a++ {
		fmt.Println("Panjang = ", satuan["panjang"])
		fmt.Println("Lebar = ", satuan["lebar"])
		fmt.Println("Tinggi = ", satuan["tinggi"])
		var luasBalok = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
		fmt.Println("Luas Balok = ", luasBalok)
	}

}
