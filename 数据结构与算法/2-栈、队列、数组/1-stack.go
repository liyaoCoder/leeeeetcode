package main

import "fmt"

/**
* @Author:liyao
* @Description:
* @Version:
* @File:1-stack
* @Date: 2022/12/23-22:04
 */
/*
	1-初始化  InitStack
	2-StackEmpty 判空
	3-Push 进栈
	4-Pop 出栈
	5-GetTop 取栈顶
	6-DestroyStack 释放栈
*/
var MAXSIZE = 50

type Stack struct {
	top  int
	size int
	data []int
}

func InitStack(size int, arr []int) *Stack {
	list := &Stack{
		size: size,
	}
	for _, ele := range arr {
		list.data = append(list.data, ele)
		list.top++
	}
	return list
}
func (s *Stack) StackEmpty() bool {
	if s.top == 0 {
		return true
	} else {
		return false
	}
}
func (s *Stack) Push(ele int) {
	if s.top == s.size {
		fmt.Println("栈已满")
		return
	}
	s.top++
	s.data = append(s.data, ele)
}
func (s *Stack) Pop() (ele int) {
	if s.top == 0 {
		fmt.Println("栈为空")
		return
	} else {
		ele = s.data[s.top-1]
		s.top--
		s.data = append(s.data[:s.top])
		return ele
	}
}
func main() {

	var arr = []int{1, 2, 3, 4}
	stack := InitStack(100, arr)
	fmt.Println("初始栈为:", stack)
	topEle := stack.Pop()
	fmt.Println("栈顶元素为：", topEle)
	fmt.Println("pop之后的stack:", stack)
	pushEle := 100
	stack.Push(pushEle)
	fmt.Println("push之后的stack:", stack)

}
