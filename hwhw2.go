package main

import "fmt"

// Стек
type Stack struct {
	stck []int
}

func (s *Stack) Add(num int) {
	s.stck = append(s.stck, num)
}
func (s *Stack) Pop() string {
	len := len(s.stck)
	if len > 0 {
		s.stck = s.stck[:len-1]
		return "выполнено"
	}
	return "пустой стек"
}
func (s *Stack) IsEmpty() bool {
	if len(s.stck) == 0 {
		return true
	}
	return false
}

// Очередь
type Queue struct {
	que []int
}

func (q *Queue) Enqueue(numb int) {
	q.que = append(q.que, numb)
}
func (q *Queue) Dequeue() string {
	len := len(q.que)
	if len > 0 {
		q.que = q.que[1:]
		return "выполнено"
	}
	return "пустая очередь"
}
func (q *Queue) qIsEmpty() bool {
	if len(q.que) == 0 {
		return true
	}
	return false
}
func main() {
	Stack := &Stack{
		stck: make([]int, 0),
	}
	fmt.Println("Стек")
	Stack.Add(1)
	Stack.Add(2)
	Stack.Add(3)
	Stack.Add(4)
	fmt.Println(Stack)
	fmt.Println(Stack.Pop())
	fmt.Println(Stack)
	fmt.Println(Stack.IsEmpty())
	fmt.Println(Stack.Pop())
	fmt.Println(Stack.Pop())
	fmt.Println(Stack.Pop())
	fmt.Println(Stack.IsEmpty())
	fmt.Println(Stack.Pop())
	fmt.Println("Очередь")
	Queue := &Queue{
		que: make([]int, 0),
	}
	Queue.Enqueue(9)
	Queue.Enqueue(8)
	Queue.Enqueue(7)
	Queue.Enqueue(6)
	fmt.Println(Queue)
	fmt.Println(Queue.Dequeue())
	fmt.Println(Queue)
	fmt.Println(Queue.qIsEmpty())
	fmt.Println(Queue.Dequeue())
	fmt.Println(Queue.Dequeue())
	fmt.Println(Queue.Dequeue())
	fmt.Println(Queue.Dequeue())
	fmt.Println(Queue.qIsEmpty())
}
