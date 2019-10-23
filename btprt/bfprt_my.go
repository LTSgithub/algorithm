package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type Bfprt struct{}

func (m *Bfprt) GetTopk(arr []int, k int) int {
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)
	return m.bfprt(copyArr, 0, len(copyArr)-1, k-1)
}

func (m *Bfprt) bfprt(arr []int, begin int, end int, i int) int {
	if begin == end {
		return arr[begin]
	}

	pivot := m.medianOfMedian(arr, begin, end)
	pivotRang := m.partition(arr, begin, end, pivot)

	if i >= pivotRang[0] && i <= pivotRang[1] {
		return arr[i]
	} else if i < pivotRang[0] {
		return m.bfprt(arr, begin, pivotRang[0]-1, i)
	} else {
		return m.bfprt(arr, pivotRang[1]+1, end, i)
	}

}

func (m Bfprt) swap(ii []int, x int, y int) {
	t := ii[x]
	ii[x] = ii[y]
	ii[y] = t
}

func (m Bfprt) partition(arr []int, l int, r int, num int) []int {
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

	return []int{less + 1, more - 1}
}

func (m *Bfprt) medianOfMedian(arr []int, begin int, end int) int {
	if begin == end {
		return arr[begin]
	}

	num := end - begin + 1

	offset := num % 5
	if offset != 0 {
		offset = 1
	}

	mArr := make([]int, num/5+offset)
	for i := 0; i < len(mArr); i++ {
		begin1 := begin + i*5
		endi := begin1 + 4
		mArr[i] = m.getMedian(arr, begin1, int(math.Min(float64(end), float64(endi))))
	}

	return m.medianOfMedian(mArr, 0, len(mArr)-1)
}

func (m *Bfprt) getMedian(arr []int, begin int, end int) int {
	copyArr := arr[begin : end+1]

	sort.Slice(copyArr, func(i, j int) bool {
		return copyArr[i] < copyArr[j]
	})

	return copyArr[len(copyArr)/2]
}

// 比较器
func comparators(arr []int, num int) {

	back := make([]int, len(arr))
	copy(back, arr)

	b := Bfprt{}
	t1 := b.GetTopk(arr, num)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	t2 := arr[num-1]

	if t1 != t2 {
		fmt.Println(back)
		fmt.Println(t1, t2)
	}

}

// 测试
func test(num int) {
	arr := []int{}
	for j := 0; j < num; j++ {
		arr = append(arr, rand.Int())
	}

	for j := 1; j <= len(arr); j++ {
		comparators(arr, j)
	}
}

func main() {

	t := 1

	for i := 0; i < t; i++ {

		for j := 1; j < 100; j++ {
			test(j)
		}
	}
}
