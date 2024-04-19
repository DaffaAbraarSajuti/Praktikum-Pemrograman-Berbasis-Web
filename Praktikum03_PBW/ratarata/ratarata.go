<<<<<<< HEAD
package ratarata

func HitungRataRata(nilaiSiswa []int) float64 {
	total := 0
	for _, nilai := range nilaiSiswa {
		total += nilai 
	}

	rataRata := float64(total) / float64(len(nilaiSiswa))
	return rataRata
=======
package ratarata

func HitungRataRata(nilaiSiswa []int) float64 {
	total := 0
	for _, nilai := range nilaiSiswa {
		total += nilai 
	}

	rataRata := float64(total) / float64(len(nilaiSiswa))
	return rataRata
>>>>>>> e375182dfa6a9d4c0f01e3182adb14aa67940038
}