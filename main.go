package main

import (
	"fmt"
)

// queueItem holds a single item in the queue.
type queueItem struct {
	next *queueItem
	data string
}

// // JSONQueue interface for json queue.
// type JSONQueue interface {
// 	getQueueSize() int
// 	insert(string) int
// 	fetch() string
// 	setWaitTime(int)
// }

// jsonQueue holds the data related to the queue.
type jsonQueue struct {
	queueHead        *queueItem
	queueTail        *queueItem
	numberOfElements int
}

func (queue *jsonQueue) getQueueSize() int {
	return queue.numberOfElements
}

func (queue *jsonQueue) insert(data string) int {
	temp := new(queueItem)
	temp.data = data
	temp.next = queue.queueHead
	queue.queueHead = temp
	if queue.queueTail == nil {
		queue.queueTail = temp
	}
	queue.numberOfElements++
	return queue.numberOfElements
}

func (queue *jsonQueue) fetch() (bool, string) {
	if queue.queueTail == nil {
		return false, ""
	}
	temp := queue.queueTail
	queue.queueTail = temp.next
	queue.numberOfElements--
	if queue.queueTail == nil {
		queue.queueHead = nil
	}
}

func NewJsonQueue() *jsonQueue {
	temp := jsonQueue{}
	temp.queueHead = nil
	temp.queueTail = nil
	temp.numberOfElements = 0
	return &temp
}

func main() {
	test := queueItem{}
	test.data = "This is a test"
	test.next = &queueItem{}
	test.next.data = "This is also a test"
	fmt.Println(test.data)
	fmt.Println(test.next.data)
}
