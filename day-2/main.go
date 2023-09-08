// peserta mampu membuat program dengan golang untuk menyelesaikan masalah.
// Soal
// 1. buatlah sebuah program untuk menentukan bilang prima

package main

import (
	"fmt"
)

// 1. buatlah sebuah program untuk menentukan bilang prima
func bilanganPrima(angka uint) bool {
	if angka <= 1 {
		return false
	}
	if angka <= 3 {
		return true
	}

	// cek bila angka bisa dibagi dengan selain 1 dan dirinya
	if angka%2 == 0 || angka%3 == 0 {
		return false
	}

	// cek faktor pembagi lain
	for i := 5; i*i <= int(angka); i += 6 {
		if int(angka)%i == 0 || int(angka)%(i+2) == 0 {
			return false
		}
	}

	return true
}

// 2. buatlah sebuah program untuk menentukan bilangan kelipatan 7
func bilanganKelipanTujuh(angka uint) bool {
	if angka < 7 {
		return false
	}

	if angka%7 == 0 {
		return true
	}

	return false
}

// 3. buatlah sebuah program untuk menghitung luas trapesium
func luasTrapesium(atas float64, bawah float64, tinggi float64) float64 {
	return ((atas + bawah) * tinggi) / 2
}

func main() {

	var angka uint = 49

	// Cek Bilangan Prima
	if bilanganPrima(angka) {
		fmt.Println(angka, " adalah bilangan prima")
	} else {
		fmt.Println(angka, " adalah bukan bilangan prima")
	}

	// Cek Bilangan kelipan 7
	if bilanganKelipanTujuh(angka) {
		fmt.Println(angka, " adalah bilangan kelipatan 7")
	} else {
		fmt.Println(angka, " adalah bukan bilangan kelipatan 7")
	}

	// luas trapesium

	atas := 2
	bawah := 2
	tinggi := 6

	fmt.Println("Luas Trapesium adalah ", luasTrapesium(float64(atas), float64(bawah), float64(tinggi)))

}
