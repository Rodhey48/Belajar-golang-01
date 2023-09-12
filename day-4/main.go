package main

import (
	"fmt"
)

// Soal no 1
type Car struct {
	types  string
	fuelIn float32
}

func (car *Car) CalculateRange() float32 {
	var result float32 = 0

	result = car.fuelIn / 1.5

	return result
}

// soal no 2

type Student struct {
	name  string
	score int
}

func CalculateStudent(data []Student) {
	var average float32
	minScore := data[0].score
	minScoreName := data[0].name
	maxScore := data[0].score
	maxScoreName := data[0].name

	tempTotalScore := 0
	for _, el := range data {
		tempTotalScore = tempTotalScore + el.score

		if minScore > el.score {
			minScore = el.score
			minScoreName = el.name
		}
		if maxScore < el.score {
			maxScore = el.score
			maxScoreName = el.name
		}
	}

	average = float32(tempTotalScore) / float32(len(data))

	fmt.Println("Average score: ", average)
	fmt.Println("Min score of student: ", minScoreName, " (", minScore, ")")
	fmt.Println("Max score of student: ", maxScoreName, " (", maxScore, ")")

}

// soal no 3
func getMinMax(number ...*int) (min int, max int) {
	var memoryMin *int = &min
	var memoryMax *int = &max

	*memoryMin = *number[0]
	*memoryMax = *number[0]

	for _, el := range number {
		if *el < *memoryMin {
			*memoryMin = *el
		}

		if *el > *memoryMax {
			*memoryMax = *el
		}
	}

	return *memoryMin, *memoryMax
}

func main() {
	fmt.Println("============ Soal 1 ===============")
	// 1. Buatlah sebuah method untuk menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang sedang terisi (1.5 L / mill), method tersebut receiver kepada struct Car yang memiliki property type dan fuelIn!
	var mobilA = Car{"sedan", 100}
	var jarak = mobilA.CalculateRange()
	fmt.Println("Jarak yang bisa ditembuh ", jarak, " mill")

	fmt.Println("============ Soal 2 ===============")

	// 	2. Buat sebuah struct dengan nama Student yang mempunyai properti name
	// dan score dalam bentuk slice kemudian simpan data siswa sebanyak 5 siswa.
	// Setelah 5 siswa dimasukkan maka program menunjukkan skor rata-rata,
	// siswa yang memiliki skor minimum dan maksimal? (implementasikan
	// method)!
	// sample test case
	// inputs:
	// inputs 1 Students name Rizky
	// inputs 1 Students score 80
	// inputs 2 Students name kobar
	// inputs 2 Students score 75
	// inputs 3 Students name Ismail
	// inputs 3 Students score 70
	// inputs 4 Students name Uman
	// inputs 4 Students score 75
	// inputs 5 Students name Yopan
	// inputs 5 Students score 60

	// Output:
	// Average Score: 72
	// Min Score of Student Yopan (60)
	// Max Score of Student Rizky (80)

	students := []Student{}
	students = append(students, Student{"Rizky", 80})
	students = append(students, Student{"Kobar", 75})
	students = append(students, Student{"Ismail", 70})
	students = append(students, Student{"Uman", 75})
	students = append(students, Student{"Yopan", 60})

	fmt.Println("Ini list student nya =>", students)
	CalculateStudent(students)

	fmt.Println("============ Soal 3 ===============")

	//  3. Buatlah program di Golang untuk menemukan nilai maksimum serta minimum
	// di antara 6 angka inputan. Gunakan multiple return fungsi, pointer untuk
	// referencing maupun deferencing!

	fmt.Println("Masukan angka:")

	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}
