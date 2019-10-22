package main

import (
	"fmt"
)

type KmpMy struct {
}

func (m *KmpMy) getNextArray(ss string) []int {
	next := []int{}

	if len(ss) == 1 {
		return []int{-1}
	}
	if len(ss) == 2 {
		return []int{-1, 0}
	}

	i := 0
	cn := 0

	for i < len(ss) {
		if ss[i-1] == ss[cn] {
			i++
			cn++
			next = append(next, cn)
		} else if cn > 0 {
			cn = next[cn]
		} else {
			i++
			next = append(next, 0)
		}
	}

	return next
}

func (m *KmpMy) getIndexof(ss string, s string) int {
	if len(ss) == 0 || len(s) == 0 || len(s) > len(ss) {
		return -1
	}

	i1 := 0
	i2 := 0
	next := m.getNextArray(s)

	for i1 < len(ss) && i2 < len(s) {
		if ss[i1] == s[i2] {
			i1++
			i2++
		} else if i2 > 0 {
			i2 = next[i2]
		} else {
			i1++
		}
	}

	if i2 == len(s) {
		return i1 - i2
	}

	return -1
}

func main() {

	kmp := KmpMy{}

	fmt.Println(kmp.getIndexof("aab", "ab"))

}
