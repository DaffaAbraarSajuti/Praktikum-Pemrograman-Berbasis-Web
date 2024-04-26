package main

import (
	"os"
	"fmt"
)

func main() {
	// Membuat direktori baru.
	err := os.Mkdir("DaffaAbraarSajuti", 0040)
	if err != nil {
	fmt.Println("Error:", err)
	return
	
	}
	fmt.Println("Direktori 'DaffaAbraarSajuti' berhasil dibuat.")
}