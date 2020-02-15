package main

import (
	"fmt"
	"time"
)

type JSONQueue interface {
	insert(string) int
	fetch() (bool, string)
	setWaitTime(int)
}

type jsonQueue struct {
	queue    []string
	waitTime int
}

func (queue *jsonQueue) insert(data string) int {
	queue.queue = append(queue.queue, data)
	return len(queue.queue)
}

func (queue *jsonQueue) fetch() (bool, string) {
	start := time.Now()
	waitTime := time.Millisecond * time.Duration(queue.waitTime)
	for (time.Since(start) < waitTime) && (len(queue.queue) == 0) {

	}
	if len(queue.queue) == 0 {
		return false, ""
	}
	temp := queue.queue[0]
	queue.queue[0] = ""
	queue.queue = queue.queue[1:]
	return true, temp
}

func (queue *jsonQueue) setWaitTime(data int) {
	queue.waitTime = data
}

func printFetch(queue *jsonQueue) {
	clear, data := queue.fetch()
	if clear {
		fmt.Print("This is the data")
		fmt.Println(data)
	} else {
		fmt.Println("Couldn not get data")
	}
}

func main() {
	temp := new(jsonQueue)
	temp.insert("data1")
	temp.insert("data2")
	temp.setWaitTime(2000)
	printFetch(temp)
	printFetch(temp)
	printFetch(temp)
}
