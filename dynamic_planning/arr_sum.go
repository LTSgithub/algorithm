package main

import (
	"fmt"
	"math"
)

const line = 5

func walk(arr [][]int, i int, j int) int {
	if i == line && j == line {
		return arr[line][line]
	}

	if i == line {
		return arr[i][j] + walk(arr, i, j+1)
	}

	if j == line {
		return arr[i][j] + walk(arr, i+1, j)
	}

	right := walk(arr, i, j+1)
	down := walk(arr, i+1, j)

	return arr[i][j] + int(math.Min(float64(right), float64(down)))

}

func main() {

	arr := [][]int{
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6},
	}

	//dp := [][]int{
	//	{26, 30, 33, 35, 36, 36},
	//	{25, 28, 30, 31, 31, 30},
	//	{24, 26, 27, 27, 26, 24},
	//	{23, 24, 24, 23, 21, 18},
	//	{22, 22, 21, 19, 16, 12},
	//	{21, 20, 18, 15, 11, 6},
	//}

	fmt.Println(walk(arr, 0, 5))

}
