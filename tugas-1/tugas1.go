package main

import (
	"errors"
	"fmt"
	"math"
)

func tugas1_2a() {
	fmt.Println("Ini contoh penggunaan 'fmt.Println'")
}

func tugas1_2b() {
	formattedString := fmt.Sprintln("Ini contoh penggunaan 'fmt.Sprintln'")
	fmt.Print((formattedString))
}

func tugas1_3a(argument float64) (float64, error) {
	// Operasi logaritma
	if argument <= 0 {
		return 0, fmt.Errorf("Ini contoh 'fmt.Errorf': Operasi log tidak bisa menggunakan argumen %f", argument)
	}
	return math.Log(float64(argument)), nil
}

func tugas1_3b(a, b float32) (float32, error) {
	// Operasi pembagian
	if b == 0 {
		return 0, errors.New("Ini contoh 'errors.New': Operasi pembagian tidak bisa menggunakan 0")
	}
	return a / b, nil
}

func main() {
	fmt.Println("")

	fmt.Println("Soal 2: Println vs Sprintln")
	tugas1_2a()
	tugas1_2b()
	fmt.Println("")

	fmt.Println("Soal 3a: fmt.Errorf")
	log100, _ := tugas1_3a(100)
	logNeg100, er := tugas1_3a(-10)
	fmt.Println(log100, logNeg100, er)

	fmt.Println("Soal 3b: errors.New")
	div2_3, _ := tugas1_3b(2, 3)
	div4_0, er := tugas1_3b(4, 0)
	fmt.Println(div2_3, div4_0, er)
}
