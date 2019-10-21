package kmp

import (
	"fmt"
)

func getNextArrayMy(arr []byte) []int {
	next := []int{}

	if len(arr) <= 1 {
		return []int{-1}
	}

	next = append(next, -1)
	next = append(next, 0)

	i := 2
	cn := 0
	for i < len(arr) {
		if arr[cn] == arr[i-1] {
			i++
			cn++
			next = append(next, cn)
		} else if cn > 0 {
			cn = next[cn]
		} else {
			i++
			next = append(next, cn)
		}
	}

	return next
}

func main() {

	fmt.Println(getNextArrayMy([]byte("aaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")))

}
