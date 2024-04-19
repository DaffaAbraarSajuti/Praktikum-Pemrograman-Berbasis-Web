package main

import "fmt"

func main() {
	fmt.Println("Bubble Sort")
	arrayNumber := [20]int{77, 31, 5, 1, 29, 66, 96, 15, 51, 37, 98, 89, 9, 22, 42, 46, 7, 13, 57, 10}
	n := len(arrayNumber)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arrayNumber[j] > arrayNumber[j+1] {
				temp := arrayNumber[j]
				arrayNumber[j] = arrayNumber[j+1]
				arrayNumber[j+1] = temp
			}
		}
	}

	fmt.Println("Setelah dilakukan pengurutan:")
	fmt.Println(arrayNumber)
}
