package main

import "fmt"

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

}
