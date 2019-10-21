package main

/*
	解决的问题：求一个数组中所有的数的左边的距离他最近的比他大的，和右边距离他最近的比他大的。
	时间复杂度：o(n)
	经典解决方法复杂度：o(n * n)
	步骤：
		1，有一个栈，这个栈由栈底到栈顶满足从大到小
		2，如果新添加的数不满足，那么弹出栈顶元素

*/

type SingleStack struct {
	CurrentIndex int // 数组下标
	CurrentValue int // 下标对应的值

	// 左
	LeftIsExist bool
	LeftIndex   int
	LeftValue   int

	// 右
	RightIsExist bool
	RightIndex   int
	RightValue   int
}

func Single(src []int) []SingleStack {
	resp := []SingleStack{}
	if len(src) == 0 {
		return resp
	}

	stackList := []SingleStack{}
	for i := 0; i < len(src); i++ {
		if len(stackList) == 0 {
			stackList = append(stackList, SingleStack{CurrentIndex: i, CurrentValue: src[i]})
			continue
		}

		//pop := stackList[len(stackList) - 1]
		if src[i] > stackList[len(stackList)-1].CurrentValue {
			stack := SingleStack{
				CurrentIndex: stackList[len(stackList)-1].CurrentIndex,
				CurrentValue: stackList[len(stackList)-1].CurrentValue,

				//LeftIsExist:
				//LeftIndex:
				//LeftValue:

				RightIsExist: true,
				RightIndex:   i,
				RightValue:   src[i],
			}

			if len(stackList) >= 2 {
				stack.LeftIsExist = true
				stack.LeftIndex = stackList[len(stackList)-1].CurrentIndex
				stack.LeftValue = stackList[len(stackList)-1].CurrentValue
			}

			resp = append(resp, stack) // 添加
			//弹出
			stackList = stackList[:len(stackList)-1]
		}

	}

	return resp
}

func main() {

}
