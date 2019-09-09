package main

import "fmt"

/*	岛问题求解：求矩阵中岛的数量  */

func isLand(arr [][]int) int {

	count := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 1 {
				infect(arr, i, j)
				count++
			}
		}
	}

	return count
}

func infect(arr [][]int, i int, j int) {

	if i < 0 || j < 0 || i >= len(arr) || j >= len(arr[0]) {
		return
	}

	if arr[i][j] != 1 {
		return
	}

	arr[i][j] = 2

	infect(arr, i+1, j)
	infect(arr, i-1, j)
	infect(arr, i, j+1)
	infect(arr, i, j-1)
}

func main() {

	arr := [][]int{
		{0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
	}

	fmt.Println(isLand(arr))

}
