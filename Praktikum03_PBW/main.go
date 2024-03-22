package main

import (
	"fmt"
	"Praktikum03_PBW/ratarata"
)

func main() {
	nilaiSiswa := []int {80, 75, 90, 85, 60}

	rataRata := ratarata.HitungRataRata(nilaiSiswa)

	fmt.Println("Rata-rata nilai siswa", rataRata)
}