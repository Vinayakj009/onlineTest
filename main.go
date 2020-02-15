package main

import "fmt"

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

func main() {
	temp := new(jsonQueue)
	temp.insert("data1")
	temp.insert("data2")
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
