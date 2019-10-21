package main

/*******************************************************************************
description: 递归获取数组最大值
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

}
