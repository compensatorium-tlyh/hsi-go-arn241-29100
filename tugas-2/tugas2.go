package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Tugas 1 //////////////////////////
func generateNik(gender string, tahun int) string {
	var nik string
	nik = "AR"
	if gender == "female" {
		nik += "T"
	} else {
		nik += "N"
	}
	nik += fmt.Sprint(tahun % 2000)

	quartile := strconv.Itoa(1 + int(time.Now().Month())/4) // 2(1)7(2)12(3)2
	idRaw := strconv.Itoa(rand.Intn(99999))
	zeros := strings.Repeat("0", 5-len(idRaw))
	nik += fmt.Sprintf("%s-%s%s", quartile, zeros, idRaw)
	// }
	return nik
}

func randomNikDataGenerator() (string, int) {
	var gender string
	if rand.Intn(10)%2 == 0 {
		gender = "male"
	} else {
		gender = "female"
	}
	tahun := 2015 + rand.Intn(10)
	return gender, tahun
}

func generateNiks(gender string, tahun int, jumlah int) []string {
	niks := make([]string, jumlah)
	for i := 0; i < jumlah; {
		gender, id := randomNikDataGenerator()
		niks[i] = generateNik(gender, id)
		i++
	}
	return niks
}

// Tugas 2 //////////////////////////
func continueNik(previousNik string, jumlah int) string {
	if len(previousNik) == 12 {
		nikSubstring := strings.Split(previousNik, "-")
		id, _ := strconv.Atoi(nikSubstring[1])
		id += 1
		zeros := strings.Repeat("0", 5-len(strconv.Itoa(id)))
		nik := nikSubstring[0] + "-" + zeros + strconv.Itoa(id)
		return nik
	} else {
		return "Error, previous NIK doesn't match the standard"
	}
}

// Tugas 3 //////////////////////////
func sortNik(niks []string) []string {
	sortedNik := make([]string, len(niks))
	copy(sortedNik, niks)
	sort.Strings(sortedNik)
	return sortedNik
}

// / Uji kode tugas
func main() {
	fmt.Println("Tugas 1:")
	fmt.Println(generateNiks("male", 2025, 5))
	fmt.Println(generateNiks("female", 2024, 1))
	fmt.Println(generateNiks("", 2019, 5))
	fmt.Println("\nTugas 2:")
	fmt.Println(continueNik("ART241-00539", 1))
	fmt.Println(continueNik("ARN191-74518", 1))
	fmt.Println(continueNik("AR191-74518", 1))
	fmt.Println("\nTugas 3:")
	// NIK telah berhasil digenerate (secara random) dan diurutkan sesuai alfabet. Adapun jika ada 2 gender maka akan terpisah berdasarkan gendernya. Pada realitanya, grup halaqah tidak digabung antara ikhwan dan akhawat
	fmt.Println(sortNik(generateNiks("", 2024, 5)))

}
