package main

import "fmt"

/*
	next数组：
		1，某一个位置前面的所有的所有的最大相等的字符串的数目
		2，如果数组小于等于1，那么next数组为{-1}
		3，如果数组等于2，那么next数组为{-1，0}
*/

// 获取next数组
func getNextArray(arrc []byte) []int {
	next := []int{}

	if len(arrc) == 1 {
		return []int{-1}
	}

	next = append(next, -1)
	next = append(next, 0)

	i := 2
	cn := 0
	for i < len(arrc) {
		if arrc[i-1] == arrc[cn] {
			cn++
			next = append(next, cn)
			i++
		} else if cn > 0 {
			cn = next[cn] //
		} else {
			next = append(next, 0)
			i++
		}
	}

	return next
}

func getIndexOf(s string, m string) int {
	if len(s) < 1 || len(m) < 1 || len(s) < len(m) {
		return -1
	}

	str1 := []byte(s)
	str2 := []byte(m)

	i1 := 0 // 代表原字符串，滑到的位置。i1会不停的往后滑动一直到结束。
	i2 := 0 // 代表子串的当前位置，是动态变化的，不是一直往后滑动，取决于next数组的值。

	next := getNextArray(str2)

	for i1 < len(str1) && i2 < len(str2) {
		if str1[i1] == str2[i2] {
			i1++
			i2++
		} else if next[i2] == -1 { // 为什么 i1++  ：因为字串的next已经走到了头，说明字串也走到了第一个字符。依然没有找到相等的，i1++要继续往后滑动来找下一个数据
			i1++
		} else {
			i2 = next[i2]
		}
	}

	if i2 == len(str2) {
		return i1 - i2
	}

	return -1
}

func main() {

	fmt.Println(getIndexOf("aaabaaaxy", "av"))

}
