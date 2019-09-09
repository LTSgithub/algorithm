package main

import (
	"fmt"
)

func sum(arr []int, i int, curSum int, aim int) bool {

	if i == len(arr) {
		return curSum == aim
	}

	return sum(arr, i+1, curSum, aim) || sum(arr, i+1, curSum+arr[i], aim)
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 8, 9, 7, 600, 5, 4, 3, 35, 7}

	fmt.Println(sum(arr, 0, 0, 600))

}
