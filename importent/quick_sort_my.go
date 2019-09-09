package main

import (
	"fmt"
	"sort"
)

type quick_sort_my struct {
}

func (m quick_sort_my) swap(ii []int, x int, y int) {
	t := ii[x]
	ii[x] = ii[y]
	ii[y] = t
}

func (m quick_sort_my) partition(arr []int, l int, r int) []int {
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

func (m quick_sort_my) quick_sort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	index := m.partition(arr, l, r)
	m.quick_sort(arr, l, index[0]-1)
	m.quick_sort(arr, index[1]+1, r)
}

func (m quick_sort_my) partition2(arr []int, l int, r int, num int) []int {
	//num := arr[r]
	less := l - 1
	more := r + 1
	index := l
	for index < more {
		if arr[index] < num {
			less++
			//m.swap(arr, less, index)
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

func (m quick_sort_my) partition1(arr []int, l int, r int, num int) []int {
	//num := arr[r]
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
		} else { // 相等
			index++
		}
	}

	return []int{less + 1, more - 1} //返回的是：左边界 右边界
}

func main() {

	sort.Slice()
	//arr := []int{1, 3, 9, 3, 7, 7, 2, 6, 8, 5, 6, 6, 3, 6, 3, 7, 7}
	arr := []int{1, 3, 9, 10, 3, 3}
	quick_sort_my{}.partition1(arr, 0, len(arr)-1, 7)
	fmt.Println(arr)

	arr1 := []int{1, 3, 9, 10, 3, 3}
	quick_sort_my{}.partition2(arr1, 0, len(arr1)-1, 7)
	fmt.Println(arr)
}
