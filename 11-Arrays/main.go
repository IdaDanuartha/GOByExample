package main

import "fmt"

func main() {
	var arrayNums [5]int
	fmt.Println("Empty:", arrayNums)

	arrayNums[4] = 100
	fmt.Println("Set:", arrayNums)
	fmt.Println("Get:", arrayNums[4])

	arrayNumsB := [5]int{1,2,3,4,5}
	fmt.Println("DCL:", arrayNumsB)

	var arrayTwoDimensions [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arrayTwoDimensions[i][j] = i + j
		}
	}

	fmt.Println("2D:", arrayTwoDimensions)
}

