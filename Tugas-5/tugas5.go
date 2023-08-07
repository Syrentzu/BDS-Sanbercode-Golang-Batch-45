package main

import (
	"fmt"
)

func main() {
	//Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
	//Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	//Soal 4

	var dataFilm = map[string]string{"judul": "", "durasi": "", "genre": "", "rilis": ""}
	// buatlah closure function disini
	var tambahDataFilm = func(judul string, durasi string, genre string, rilis string) {
		for i := 0; i < 1; i++ {
			dataFilm["judul"] = judul
			dataFilm["durasi"] = durasi
			dataFilm["genre"] = genre
			dataFilm["rilis"] = rilis
			fmt.Println(dataFilm)
		}

	}
	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

// Soal 1
func luasPersegiPanjang(panjang int, lebar int) int {
	var luasPersegi = panjang * lebar
	return luasPersegi
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	var kelilingPersegi = 2 * (panjang + lebar)
	return kelilingPersegi
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	var volumeB = panjang * lebar * tinggi
	return volumeB
}

// Soal 2
func introduce(Nama string, Gender string, Pekerjaan string, Umur string) (Kalimat string) {
	if Gender == "laki-laki" {
		Kalimat = "Pak " + Nama + " adalah seorang " + Pekerjaan + " yang berusia " + Umur + " tahun"
		return Kalimat
	} else {
		Kalimat = "Ibu " + Nama + " adalah seorang " + Pekerjaan + " yang berusia " + Umur + " tahun"
		return Kalimat
	}
}

// Soal 3
func buahFavorit(Nama string, Input2 ...string) (Kalimat1 string) {
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"
	var tampung = ""
	for _, Input := range Input2 {
		tampung += Input
	}
	Kalimat1 = "halo nama saya " + Nama + " dan buah favorit saya adalah " + tampung

	return Kalimat1
}
