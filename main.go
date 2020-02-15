package main

import (
	"fmt"
	"time"
)

type JSONQueue interface {
	insert(string) int
	fetch() (bool, string)
	setWaitTime(time.Duration)
}

type jsonQueue struct {
	queue    []string
	waitTime time.Duration
}

func (queue *jsonQueue) insert(data string) int {
	queue.queue = append(queue.queue, data)
	return len(queue.queue)
}

func (queue *jsonQueue) fetch() (bool, string) {
	start := time.Now()
	for (time.Since(start) < queue.waitTime) || (len(queue.queue) == 0) {

	}
	if len(queue.queue) == 0 {
		return false, ""
	}
	temp := queue.queue[0]
	queue.queue[0] = ""
	queue.queue = queue.queue[1:]
	return true, temp
}

func (queue *jsonQueue) setWaitTime(data time.Duration) {
	queue.waitTime = data
}

func main() {
	temp := new(jsonQueue)
	start := time.Now()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	temp.insert("data1")
	temp.insert("data2")
	temp.setWaitTime(time.Second * 4)
	clear, temp1 := temp.fetch()
	if clear {
		fmt.Print("This is the first data ")
		fmt.Println(temp1)
	} else {
		fmt.Println("Could not fetch 1st data")
		return
	}
	clear, temp1 = temp.fetch()
	if clear {
		fmt.Print("This is the 2nd data ")
		fmt.Println(temp1)
	} else {
		fmt.Println("Could not fetch 2nd data")
		return
	}
	clear, temp1 = temp.fetch()
	if clear {
		fmt.Print("This is the 3rd data ")
		fmt.Println(temp1)
	} else {
		fmt.Println("Could not fetch 3rd data")
		return
	}
}
