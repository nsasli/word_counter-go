package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Aplikasi Penghitung Kata")

	// Meminta pengguna untuk memasukkan teks
	var teks string
	fmt.Print("Masukkan teks: ")
	fmt.Scanln(&teks)

	// Memisahkan teks menjadi kata-kata
	kataKata := strings.Fields(teks)

	// Membuat peta untuk menghitung jumlah kemunculan setiap kata
	kemunculanKata := make(map[string]int)

	// Menghitung jumlah kemunculan setiap kata
	for _, kata := range kataKata {
		kemunculanKata[kata]++
	}

	// Menampilkan hasil penghitungan
	fmt.Println("Jumlah kemunculan kata dalam teks:")
	for kata, jumlah := range kemunculanKata {
		fmt.Printf("%s: %d\n", kata, jumlah)
	}
}
