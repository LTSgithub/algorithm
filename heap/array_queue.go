package main

import "fmt"

type ArrayQueue struct {
	ArrAy []int // 数据
	Start int   // 指向队列的开头
}

func (m *ArrayQueue) Push(data int) {

	m.ArrAy = append(m.ArrAy, data)

	return
}

func (m *ArrayQueue) Pop() (int, error) {

	if len(m.ArrAy)-m.Start <= 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	out := m.ArrAy[m.Start]
	m.Start++

	return out, nil
}

func main() {

	queue := ArrayQueue{}

	for i := 0; i < 10; i++ {
		queue.Push(i)
	}

	for i := 0; i < 15; i++ {
		fmt.Println(queue.Pop())
	}

}
