package main

import "fmt"

/**
* @Author:liyao
* @Description:
* @Version:
* @File:main
* @Date: 2022/12/22-21:19
 */
/*线性表----数组*/
// 插入-删除-按值查找
var MaxSize = 50

type L struct {
	Length  int
	MaxSize int
	Data    []int
}

// 1-初始化

func InitSqList(data []int, maxSize int) *L {
	list := &L{
		Length:  len(data),
		Data:    data,
		MaxSize: maxSize,
	}
	return list
}

// 2-查找

func Search(data int, list *L) int {
	for i, element := range list.Data {
		if element == data {
			return i + 1
		}
	}
	return -1
}

//3- 插入

func InsertList(element int, list *L, index int) {
	if index >= list.MaxSize {
		return
	}
	if list.MaxSize == list.Length {
		return
	}
	var ele int
	list.Data = append(list.Data, ele)
	fmt.Println("当前长度为:", len(list.Data))
	for x := list.Length; x >= index-1; x-- {
		list.Data[x] = list.Data[x-1]
	}
	list.Data[index-1] = element
	list.Length++
}

// 4-删除
func DeleteEle(index int, l *L) {
	if index < 0 || index > l.Length {
		fmt.Println("越界")
		return
	}
	for i := index; i < l.Length; i++ {
		l.Data[i-1] = l.Data[i]
	}
	l.Length--
	l.Data = append(l.Data[:l.Length])
}
func main() {
	var data = []int{1, 2, 3, 4, 5}
	maxSize := 50
	list := InitSqList(data, maxSize)
	fmt.Println("初始list为:", list)
	InsertList(100, list, 3)
	fmt.Println("插入后的list:", list)
	index := Search(100, list)
	if index == -1 {
		fmt.Println("list中无该数据")
	} else {
		fmt.Println("该数据位置为:", index)
	}
	DeleteEle(1, list)
	fmt.Println("删除第一个数据后的list为：", list)

}
