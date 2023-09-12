package main

import (
	featurea "belajarGolang/featureA" // go mod
	"errors"
	"fmt"
	"slices"
)

func main() {
	age := 9 // sama aja dengan var age int = 9
	fmt.Println(age)

	var umur int = 33

	// bitwise (>> << ) pergeseran, mengubah biner dan digeser di tambah 0 => 10(dalam biner) menjadi 100

	// penkondisian
	if age >= 17 || umur >= 8 {
		// exekusi
		fmt.Println("masuk bos", age)
	} else {
		// exekusi
		fmt.Println(umur)
	}

	// perulangan
	for i := 0; i < age; i++ {
		fmt.Println("print ke ", i)
	}

	// 	Overview
	// peserta mampu membuat program dengan golang untuk menyelesaikan masalah.
	// Soal
	// 1. buatlah sebuah program untuk menentukan bilang prima
	// 2. buatlah sebuah program untuk menentukan bilangan kelipatan 7
	// 3. buatlah sebuah program untuk menghitung luas trapesium

	// type data
	// array
	var array [5]int = [5]int{1, 2, 3, 4, 5} // sudah dipesan array ukuran 5 index
	fmt.Println(array)
	//  array bersifat static, index sudah ditetapkan, gak bisa dihapus dan ditambah, nilai bila tidak diisi akan berisi 0, ukuruan statik

	// slice
	// sama kayak array tapi semua dinamis, bisa ditambah atau dikurangi
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice[1] = 3
	fmt.Println(slice)

	// delete
	slice = slices.Delete(slice, 2, 3)
	fmt.Println(slice, "after delete")

	// type data MAP
	// menyimpan banyak nilai, berbasis key, ukuran dinamis

	var pendapatan = map[string]int{} // key string, prperti int
	pendapatan["januari"] = 1000
	pendapatan["februari"] = 2000
	pendapatan["maret"] = 3000

	fmt.Println(pendapatan)
	fmt.Println(pendapatan["maret"])

	result, isFounf := pendapatan["april"] // kelibihanya klo tidak ada bisa didapat value nya

	fmt.Println(result, isFounf)

	// Pointer
	// adalah alamat memory di komp

	var poin int = 20
	fmt.Println(&poin)            // akan muncul alamat memory nya
	var alamatMemory *int = &poin // fungsi bintang untuk deklarasi int pointer berisi alamat
	fmt.Println((alamatMemory))

	// salah satu fungsi pointer untuk merubah variable dengan menuju lngsg ke alamatnya
	*alamatMemory = 30
	// lebih effective dari pada mengirm banyak data

	// bentuk struktur data example object

	var rudi User

	// interface

	// var rudiInterface UserInterface
	rudi.age = 65
	rudi.nama = "rudi"
	rudi.email = "email@hhh.ll"

	fmt.Println(rudi)
	fmt.Println(rudi.email)
	// method
	// mirip funct di object
	rudi.ChangeNameEmail()

	// 	rudiInterface = rudi

	// rudiInterface.ChangeNameEmail()

	// import module dari folder lain

	// buat g mod dulu => go mod init belajarGolang
	featurea.FeatureA()

	//error handle

	resresult, err := tambah(1, 3)
	fmt.Println(resresult, err)

	// panic
	// untuk memberhentikan proses
	// panic("ini panic")

	// recover
	// untuk menangkap error like try catch

}

// untuk watch di terminal
// nodemon --exec go run main.go --signal SIGTERM

type User struct {
	nama  string
	age   int
	email string
}

// method
// mirip funct di object

func (rudi *User) ChangeNameEmail() { // jngan lupa kasiih biintang
	rudi.nama = "blalalala"
}

// interface
// menghubungkan fungc atau daftar method struct

type UserInterface interface {
	ChangeNameEmail()
}

// type yang dinamis like any => interface{}
func show(data interface{}) {
	switch data.(type) {
	case int:
		fmt.Println(data)

	case string:
		fmt.Println(data)
	}
}

func tambah(a, b int) (int, error) {
	result := a + b
	return result, errors.New("Errororor")
}
