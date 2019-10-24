package main

import "fmt"

/*******************************************************************************
ps:
	1，求数组内最大值
	2，使用递归
*********************************************************************************/
func getMax(arr []int, l int, r int) int {

	if l == r {
		return arr[l]
	}

	mid := (l + r) / 2
	x := getMax(arr, l, mid)
	y := getMax(arr, mid+1, r)
	if x > y {
		return x
	}

	return y
}

func main() {

	arr := []int{1, 2, 3, 4, 56, 7, 89, 100}

	fmt.Println(getMax(arr, 0, len(arr)-1))

}
