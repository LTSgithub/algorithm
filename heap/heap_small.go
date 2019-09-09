package main

import "fmt"

// 大到小

type heap_small struct {
}

func (m heap_small) swap(arr []int, x int, y int) {
	arr[x], arr[y] = arr[y], arr[x]
}

func (m heap_small) insert(arr []int, index int) []int {

	for arr[index] < arr[(index-1)/2] {
		m.swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}

	return arr
}

// 某个位置上面数变大，往下沉淀
func (m heap_small) heapify(arr []int, index int, size int) {

	left := index*2 + 1

	for left < size {
		smallLeast := left

		if left+1 < size && arr[left+1] < arr[left] {
			smallLeast = left + 1
		}
		if arr[index] < arr[smallLeast] {
			smallLeast = index
		}
		if index == smallLeast {
			break
		}

		m.swap(arr, index, smallLeast)
		index = smallLeast
		left = index*2 + 1
	}

	return
}

func (m heap_small) sort_big_to_small(arr []int) {

	for i := 0; i < len(arr); i++ {
		m.insert(arr, i)
	}

	size := len(arr)

	for size > 0 {
		size--
		m.swap(arr, 0, size)
		m.heapify(arr, 0, size)
	}

}

func main() {

	h := heap_small{}
	arr := []int{10, 2, 3, 4, 5, 6, 7, 7, 8, 990}

	fmt.Println(arr)
	h.sort_big_to_small(arr)
	fmt.Println(arr)
}
