package main

import "fmt"

func getNextArrayMy(arr []byte) []int {
	next := []int{}

	if len(arr) == 1 {
		return []int{-1}
	}

	next = append(next, -1)
	next = append(next, 0)

	index := 0
	current := 2
	for current < len(arr) {
		if arr[index] == arr[current] {
			current++
			index++
			next = append(next, index)
		} else if index > 0 {

		} else {
			current++
		}

	}

	return next
}

func getIndexOfMy(s string, m string) int {

	return -1
}

func main() {

	fmt.Println(getIndexOfMy("aaaaaaaaaaavvvvvvvvvvvvvv", "av"))

}
