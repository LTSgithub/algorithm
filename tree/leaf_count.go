package tree

import "fmt"

type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

func prePrint1(node *Node) int {

	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 1
	}

	a := prePrint1(node.Left)
	b := prePrint1(node.Right)

	return a + b
}

func main() {

	node_1_1 := &Node{Data: "a"}

	node_2_1 := &Node{Data: "b"}
	node_2_2 := &Node{Data: "c"}
	{
		node_1_1.Left = node_2_1
		node_1_1.Right = node_2_2
	}

	node_3_1 := &Node{Data: "d"}
	node_3_2 := &Node{Data: "e"}
	node_3_3 := &Node{Data: "f"}
	node_3_4 := &Node{Data: "g"}
	{
		node_2_1.Left = node_3_1
		node_2_1.Right = node_3_2
		node_2_2.Left = node_3_3
		node_2_2.Right = node_3_4
	}

	node_4_1 := &Node{Data: "h"}
	node_4_2 := &Node{Data: "i"}
	node_4_3 := &Node{Data: "j"}
	{
		node_3_1.Left = node_4_1
		node_3_1.Right = node_4_2
		node_3_2.Left = node_4_3
	}

	prePrint1(node_1_1)

	fmt.Println(prePrint1(node_1_1))

}
