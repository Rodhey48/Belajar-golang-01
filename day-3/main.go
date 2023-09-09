package main

import (
	"fmt"
	"strconv"
	"strings"
)

// soal no 1
func ArrayMerge(arrayA, arrayB []string) []string {
	if len(arrayA) > 0 && len(arrayB) == 0 {
		return arrayA
	}

	if len(arrayA) == 0 && len(arrayB) > 0 {
		return arrayB
	}

	if len(arrayA) == 0 && len(arrayB) == 0 {
		return arrayA
	}

	var result = arrayA
	if len(arrayA) > 0 && len(arrayB) > 0 {
		for _, el := range arrayB {
			flag := true
			for _, val := range result {
				if el == val {
					flag = false
				}
			}
			if flag {
				result = append(result, el)
			}
		}
	}
	return result
}

// soal no 2
func Mapping(slice []string) map[string]int {
	result := map[string]int{}
	if len(slice) > 0 {
		for _, el := range slice {
			_, isFound := result[el]
			if !isFound {
				result[el] = 1
			} else {
				result[el]++
			}
		}
	}
	return result
}

// soal no 3
func munculSekali(angka string) []int {
	result := []int{}
	splited := strings.Split(angka, "") // di split biar gampang di looping

	tempCount := make(map[string]int) // wadah untuk menampung jumlah nilai
	if len(splited) > 0 {
		for _, el := range splited {
			tempCount[el]++
		}
		for i, val := range tempCount { // kalo looping type map i (key) value (nilai nya)
			str, _ := strconv.Atoi(i) // merubah string ke int karena result type int
			if val == 1 {
				result = append(result, str)
			}
		}
	}
	return result
}

func main() {
	fmt.Println("============ Soal 1 ===============")
	// 1. Buatlah sebuah program menggabungkan 2 array yang diberikan, dan jangan sampai terdapat nama yang sama di data yang sudah tergabung tadi!
	// Test case

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergie", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergie", "jin", "steve", "bryan",]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []

	fmt.Println("============ Soal 2 ===============")
	// 2. buatlah sebuah program yang dapat menghitung berapa banyak sebuah string yang sama didalam sebuah slice!
	// Test case
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	//map[adi:1 asd:2 qwe:3]

	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	//map[asd:2 qwe:1]

	fmt.Println(Mapping([]string{}))
	//map[]

	fmt.Println("============ Soal 3 ===============")
	// 3. Buat program sesuai dengan deskripsi di bawah. Input merupakan variable string berisi kumpulan angka. Output merupakan list / array berisi angka yang hanya muncul 1 kali pada input
	// Test case
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [ 1 2 3 4 4]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // 8 7 2 5 4

}
