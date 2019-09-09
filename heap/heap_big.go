package main

import "fmt"

type heap_big struct {
}

func (m heap_big) swap(ii []int, x int, y int) {
	ii[x], ii[y] = ii[y], ii[x]
}

// 形成大根堆
func (m heap_big) insert(ii []int, index int) []int {

	for ii[(index-1)/2] < ii[index] {
		m.swap(ii, (index-1)/2, index)
		index = (index - 1) / 2
	}

	return ii
}

//某个数变小，往下沉淀
func (m heap_big) heapify(ii []int, index int, size int) {

	left := index*2 + 1

	for left < size {
		largest := left
		if left+1 < size && ii[left+1] > ii[left] {
			largest = left + 1
		}

		if ii[index] > ii[largest] {
			largest = index
		}

		if index == largest {
			break
		}

		m.swap(ii, index, largest)
		index = largest
		left = largest*2 + 1
	}

	return
}

// 从小到大排序
func (m heap_big) sort_small_to_big(arr []int) {

	for i := 0; i < len(arr); i++ {
		m.insert(arr, i) //时间复杂度：log1 + log2 + log3 + logn = n
	}

	size := len(arr)
	for size > 0 {
		m.heapify(arr, 0, size)
		size--
		m.swap(arr, 0, size)
	}
}

func main() {

	h := heap_big{}
	arr := []int{10, 2, 3, 4, 5, 6, 7, 7, 8, 990}

	fmt.Println(arr)
	h.sort_small_to_big(arr)
	fmt.Println(arr)

}
