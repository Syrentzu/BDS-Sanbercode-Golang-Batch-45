package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var panjangPersegiPanjangCon, _ = strconv.Atoi(panjangPersegiPanjang)
	var lebarPersegiPanjangCon, _ = strconv.Atoi(lebarPersegiPanjang)

	var alasSegitigaCon, _ = strconv.Atoi(alasSegitiga)
	var tinggiSegitigaCon, _ = strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjangPersegiPanjangCon * lebarPersegiPanjangCon
	var kelilingPersegiPanjang int = 2 * (panjangPersegiPanjangCon + lebarPersegiPanjangCon)
	var luasSegitiga int = alasSegitigaCon * tinggiSegitigaCon / 2

	fmt.Println(luasPersegiPanjang)
	fmt.Println(kelilingPersegiPanjang)
	fmt.Println(luasSegitiga)

	//Soal 2
	var nilaiJohn = 80
	// var nilaiDoe = 50
	var PredikatNilai string

	if nilaiJohn >= 80 {
		PredikatNilai = "A"
		fmt.Println(PredikatNilai)
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		PredikatNilai = "B"
		fmt.Println(PredikatNilai)
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		PredikatNilai = "C"
		fmt.Println(PredikatNilai)
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		PredikatNilai = "D"
		fmt.Println(PredikatNilai)
	} else {
		PredikatNilai = "E"
		fmt.Println(PredikatNilai)
	}

	//Soal 3
	var tanggal = 26
	var bulan = 10
	var tahun = 2004

	var tanggalCon = strconv.Itoa(tanggal)
	var bulanCon = strconv.Itoa(bulan)
	var tahunCon = strconv.Itoa(tahun)

	switch bulan {
	case 1:
		bulanCon = "Januari"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)

	case 2:
		bulanCon = "Februari"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)

	case 3:
		bulanCon = "Maret"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)

	case 4:
		bulanCon = "April"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)

	case 5:
		bulanCon = "Mei"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)
	case 6:
		bulanCon = "Juni"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)

	case 7:
		bulanCon = "Juli"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)
	case 8:
		bulanCon = "Agustus"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)

	case 9:
		bulanCon = "September"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)

	case 10:
		bulanCon = "Oktober"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)

	case 11:
		bulanCon = "November"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)

	case 12:
		bulanCon = "Desember"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon
		fmt.Println(halo)

	}

	//Soal 4
	if tahun >= 1944 && tahun <= 1964 {
		var Generasi = "Baby boomer"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon + " Merupakan " + Generasi
		fmt.Println(halo)
	} else if tahun >= 1965 && tahun <= 1979 {
		var Generasi = "Generasi X"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon + " Merupakan " + Generasi
		fmt.Println(halo)
	} else if tahun >= 1980 && tahun <= 1994 {
		var Generasi = "Generasi Y (Millenials)"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon + " Merupakan " + Generasi
		fmt.Println(halo)
	} else if tahun >= 1995 && tahun <= 2015 {
		var Generasi = "Generasi Z"
		var halo = tanggalCon + " " + bulanCon + " " + tahunCon + " Merupakan " + Generasi
		fmt.Println(halo)
	}

}
