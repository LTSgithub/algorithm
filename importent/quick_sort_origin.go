package main

import (
	"fmt"
)

type quick_sort_origin struct {
}

func (m quick_sort_origin) swap(ii []int, x int, y int) {
	t := ii[x]
	ii[x] = ii[y]
	ii[y] = t
}

func (m quick_sort_origin) partition(arr []int, l int, r int) []int {
	num := arr[r]
	less := l - 1
	more := r + 1
	index := l
	for index < more {
		if arr[index] < num {
			less++
			m.swap(arr, less, index)
			index++
		} else if arr[index] > num {
			more--
			m.swap(arr, more, index)
		} else {
			index++
		}
	}

	return []int{less + 1, more - 1} //返回的是：左边界 右边界
}

func (m quick_sort_origin) quick_sort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	index := m.partition(arr, l, r)
	m.quick_sort(arr, l, index[0]-1)
	m.quick_sort(arr, index[1]+1, r)
}

func main() {

	arr := []int{1, 3, 9, 3, 7, 7, 2, 6, 8, 5, 6, 6, 3, 6, 3, 7, 7}

	fmt.Println(arr)
	quick_sort_origin{}.quick_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
