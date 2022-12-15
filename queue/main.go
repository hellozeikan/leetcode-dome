package main

import "fmt"

type Queue struct {
	arr  []int
	size int
}

// 数组向前面出
func (que *Queue) Pop() int {
	if que.size == 0 {
		return -1
	}
	que.size--
	temp := que.arr[0]
	que.arr = que.arr[1:]
	return temp
}
func (que *Queue) Push(val int) {
	que.arr = append(que.arr, val)
	que.size++
}

// empty
// size
func (que *Queue) Empty() bool {
	if que.size > 0 {
		return false
	}
	return true
}

func (que *Queue) Size() int {
	return que.size
}
func main() {
	queue := Queue{[]int{}, 0}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	fmt.Println(queue.Size())
	fmt.Println(queue.Pop())
	queue.Push(5)
	fmt.Println(queue.Pop())
}

// 延时队列
