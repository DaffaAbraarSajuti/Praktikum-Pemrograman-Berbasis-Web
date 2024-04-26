package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	err = os.Chmod("DaffaAbraarSajuti", 0040)
	if err != nil{
		fmt.Println("Error:", err )
		return
	}
	fmt.Println("Izin 'DaffaAbraarSajuti' telah diubah menjadi 0040")
}