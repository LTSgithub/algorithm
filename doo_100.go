package main

import (
	"fmt"
	"strings"
)

// 判断一个整数是不是回文
func doo_0009(num int) bool {

	if num < 0 || (num != 0 && num%10 == 0) {
		return false
	}

	rever := 0
	for num > rever {
		rever = rever*10 + num%10
		num = num / 10
	}

	return num == rever || num == rever/10
}

// 这个是按照自己的思路写的，狗屎啊，但是能用
func doo_1047_1(ss string) string {

	start := 0
	end := 0

	stack := []byte{}

	for i := 0; i < len(ss); {
		if len(stack) < 1 {
			start++
			end++
			stack = append(stack, ss[i])
			i++
		} else if stack[len(stack)-1] != ss[i] {
			if start == end {
				stack = append(stack, ss[i])
				start++
				end++
				i++
			} else {
				aa := strings.TrimRight(string(stack), string(stack[len(stack)-1]))
				stack = []byte(aa)
				if len(stack) == 0 {
					start = 0
					end = 0
				} else {
					start = len(stack) - 1
					end = len(stack) - 1
				}
			}
		} else if stack[len(stack)-1] == ss[i] {
			end++
			stack = append(stack, ss[i])
			i++
		}
	}

	if len(stack) >= 2 && stack[len(stack)-1] == stack[len(stack)-2] {
		return strings.TrimRight(string(stack), string(stack[len(stack)-2]))
	}

	return string(stack)
}

func main() {

	fmt.Println(doo_1047("1221"))

}
