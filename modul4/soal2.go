package main

import	"fmt"

func hitungSkor() (int, int) {
	var soal, skor, waktu int

	soal = 0
	skor = 0

	for i := 0; i< 8; i++ {
		fmt.Scan(&waktu)

		if waktu <= 300 {
			soal = soal + 1
			skor = skor + waktu
		}
	}
	return soal, skor
	
}

func main() {
	var nama string
	var pemenang string
	var maxSoal, minSkor int
	var soal, skor int

	maxSoal = -1
	minSkor = 999999

	fmt.Scan(&nama)

	for nama != "Selesai"{

		soal, skor = hitungSkor()
		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
		fmt.Scan(&nama)
	}
	fmt.Println(pemenang, maxSoal, minSkor)
}