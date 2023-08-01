package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//Soal 1
	var kata1 = "Bootcamp"
	var kata2 = "Digital"
	var kata3 = "Skill"
	var kata4 = "Sanbercode"
	var kata5 = "Golang"

	fmt.Println(kata1, kata2, kata3, kata4, kata5)

	//Soal 2
	halo := "Halo Dunia"

	var tampung = strings.Replace(halo, "Dunia", "Golang", 6)

	fmt.Println(tampung)

	//Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	var lengkap1 = kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat

	fmt.Println(lengkap1)

	//Soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var convPertama, err1 = strconv.Atoi(angkaPertama)
	var convKedua, err2 = strconv.Atoi(angkaKedua)
	var convKetiga, err3 = strconv.Atoi(angkaKetiga)
	var convKeempat, err4 = strconv.Atoi(angkaKeempat)

	fmt.Println(convPertama, ",", err1)
	fmt.Println(convKedua, ",", err2)
	fmt.Println(convKetiga, ",", err3)
	fmt.Println(convKeempat, ",", err4)

	//Soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	var convert1 = strings.Replace(kalimat, "halo", "Hi", 2)
	var convert2 = strconv.Itoa(angka)

	fmt.Println("\"" + convert1 + "\"" + " - " + convert2)
}
