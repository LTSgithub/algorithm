package main

import "fmt"

type ArrayStack struct {
	Arr   []int
	Index int
}

func (m *ArrayStack) Lenth() int {
	return len(m.Arr)
}

func (m *ArrayStack) Push(data int) {
	m.Arr = append(m.Arr, data)
	m.Index++
}

func (m *ArrayStack) Pop() (int, error) {

	if m.Index < 1 {
		return 0, fmt.Errorf("stack is empty")
	}

	out := m.Arr[m.Index-1]
	m.Index--

	return out, nil
}

func main() {
	stack := ArrayStack{}

	for i := 0; i < 100; i++ {
		stack.Push(i)
	}

	for i := 0; i < stack.Lenth(); i++ {
		fmt.Println(stack.Pop())
	}

}
