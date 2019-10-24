package main

import (
	"container/list"
	"fmt"
)

/*
	题目：图片01——1
	一个大小为3的窗口划过一个数组，求所有的窗口内的最大值
*/

func GetMaxWindow(arr []int, size int) []int {
	resp := []int{}

	ls := list.New() // 存放数组的下标

	for i := 0; i < len(arr); i++ {
		for ls.Len() != 0 && arr[ls.Back().Value.(int)] <= arr[i] {
			ls.Remove(ls.Back())
		}

		ls.PushBack(i)

		if ls.Front().Value.(int) == i-size {
			ls.Remove(ls.Front())
		}

		if i >= size-1 {
			resp = append(resp, ls.Front().Value.(int))
		}
	}

	return resp
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	indexList := GetMaxWindow(arr, 2) //返回的是下标

	for _, v := range indexList {
		fmt.Print(" ", arr[v])
	}

}
