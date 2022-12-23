package main

import (
	"fmt"
)

/**
* @Author:liyao
* @Description:
* @Version:
* @File:
* @Date: 2022/12/23-21:09
 */

// 单链表结构体
type Node struct {
	data interface{}
	next *Node
}
type List struct {
	length   int
	headNode *Node
}

// 1-初始化
func InitList() *List {
	var node *Node
	var list List
	list.headNode = node
	return &list
}

// 2-插入
func (l *List) InsertNode(index int, node *Node) {
	if index <= 0 || index > l.length {
		fmt.Println("our of range")
		return
	}
	// 在第index个位置插入，就是在索引为0开始index-1处插入
	pre := l.headNode
	if index == 1 {
		node.next = pre
		l.headNode = node
	} else {
		for i := 1; i < index-1; i++ {
			pre = pre.next
		}
		node.next = pre.next
		pre.next = node
	}
	l.length++
}

// 3-删除指定元素
func (l *List) DeleteNode(v interface{}) {
	pre := l.headNode
	for i := 1; i < l.length; i++ {
		prepre := pre
		pre = pre.next
		if pre.data == v {
			prepre.next = pre.next
		}
	}
}

//4-单链表查值 ------------------直接遍历即可

func main() {
	list := InitList()
	node := &Node{
		data: 1,
		next: nil,
	}
	list.headNode = node
	list.length++
	fmt.Println("初始list为:", list)
	insertNode := &Node{
		data: 2,
		next: nil,
	}
	list.InsertNode(1, insertNode)
	fmt.Println("插入元素后的list head:", list.headNode)
	fmt.Println("list head的下一个元素:", list.headNode.next)
	fmt.Println("list length：", list.length)
}
