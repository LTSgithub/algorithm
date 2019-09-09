package main

import "fmt"

func swap(ii []int, x int, y int) {
	t := ii[x]
	ii[x] = ii[y]
	ii[y] = t
}

func heapInsert(arr []int, index int, size int) []int {
	if index > size-1 || size > len(arr) {
		return arr
	}

	for arr[(index-1)/2] > arr[index] {
		swap(arr, (index-1)/2, index)
		index = (index - 1) / 2
	}

	return arr
}

func heapify(arr []int, index int, size int) {

	if index > size-1 || size > len(arr) {
		return
	}
	left := index*2 + 1
	for left < size {
		least := left
		if arr[left] > arr[left+1] {
			least = left + 1
		}
		if arr[least] > arr[index] {
			least = index
		}
		if least == index {
			break
		}
		left = index*2 + 1
		swap(arr, index, least)
	}

}

func doo_217(arr []int, k int) int {

	if len(arr) < k || k <= 0 {
		return -1
	}

	t := []int{}
	for i := 0; i < len(arr); i++ {
		if len(t) < k {
			t = append(t, arr[i])
			t = heapInsert(t, len(t)-1, len(t))
		} else if arr[i] > t[0] {
			t[0] = arr[i]
			heapify(t, 0, len(t))
		}
	}

	return t[0]
}

func main() {

	fmt.Println(doo_217([]int{10, 9, 8, 8, 7, 6, 5, 4, 3, 2, 1}, 4))

}
