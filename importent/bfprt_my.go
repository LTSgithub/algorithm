package main

type prt_my struct{}

func (m *prt_my) bfprt(arr []int, begin int, end int, k int) int {

	modian := m.medianOfData()
	equalRang := m.partition()

	if k >= equalRang[0] && k <= equalRang[1] {
		return arr[k]
	} else if k < equalRang[0] {
		return m.bfprt(arr, begin, equalRang[0]-1, k)
	} else {
		return m.bfprt(arr, equalRang[1]+1, end, k)
	}
}

func (m *prt_my) medianOfData() int {

}

func (m *prt_my) swap(ii []int, x int, y int) {
	ii[x], ii[y] = ii[y], ii[x]
}

func (m *prt_my) partition(arr []int, l int, r int, num int) []int {

	less := l - 1
	more := r + 1
	index := 1

	for index < more {
		if arr[index] < num {
			index++
			less++
		} else if arr[index] == num {

		} else {
			m.swap(arr, index, more)
			less--
		}

	}

	for index < more {
		if arr[index] < num {
			less++
			swap(arr, less, index)
			index++
		} else if arr[index] > num {
			more--
			swap(arr, more, index)
		} else {
			index++
		}
	}

}

func (m *prt_my) getMinK(arr []int, k int) {
	back := arr[:len(arr)-1]

	m.bfprt(back, 0, len(back)-1, k-1)
}

func main() {

}
