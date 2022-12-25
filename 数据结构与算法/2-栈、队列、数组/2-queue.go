package main

import "fmt"

/**
* @Author:liyao
* @Description:
* @Version:
* @File:2-queue
* @Date: 2022/12/23-23:25
 */
/*
	1-InitQueue
	2-QueueEmpty
	3-EnQueue  入队
	4-DeQueue 出队
	5-GetHead 队头
*/

type Queue struct {
	data    []int
	front   int
	rear    int
	maxSize int
}

func InitQueue() Queue {
	var queue Queue
	//var data [50]int
	//queue.data = data
	queue.front = 0
	queue.rear = 0
	queue.maxSize = 50
	return queue
}
func (q *Queue) QueryEmpty() bool {
	if q.front == q.rear && q.front == 0 {
		return true
	}
	return false
}
func (q *Queue) EnQueue(ele int) bool {
	if q.rear+1 == q.maxSize {
		fmt.Println("队列已满")
		return false

	}
	q.rear++
	q.data = append(q.data, ele)

	return true
}
func (q *Queue) DeQueue() (ele int) {
	if q.rear == q.front && q.rear == 0 {
		fmt.Println("队列为空")
		return -1
	}
	ele = q.data[q.front]
	q.data = q.data[1:]
	q.front = (q.front + 1) % q.maxSize
	return ele
}
func (q *Queue) GetHead() int {
	if q.front == q.rear {
		fmt.Println("队空")
		return -1
	}
	ele := q.data[q.front]
	return ele
}
func main() {
	queue := InitQueue()
	fmt.Println("初始队列：", queue)
	if queue.QueryEmpty() {
		fmt.Println("队列为空")
	}
	if queue.EnQueue(100) && queue.EnQueue(200) {
		fmt.Println("入队成功")
		fmt.Println("当前队列为:", queue)
	}

	head := queue.GetHead()
	fmt.Println("队头:", head)
	ele := queue.DeQueue()
	fmt.Println("ele:", ele)
	fmt.Println("当前队列：", queue)
}
