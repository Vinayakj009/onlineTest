package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func (queue *jsonQueue) insert(data string) int {
	if !isJSON(data) {
		return len(queue.queue)
	}
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

var StorageQueue = new(jsonQueue)

func handler(responseWriter http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	if request.URL.Path[1:] == "fetch" {
		clear, data := StorageQueue.fetch()
		if data == "" {
			data = "\"\""
		}
		fmt.Fprintf(responseWriter, `{"success":%t,"data":%s}`, clear, data)
	} else if request.URL.Path[1:] == "insert" {
		data := request.Form.Get("data")
		if isJSON(data) == false {
			fmt.Fprintf(responseWriter, `{"success":%t}`, false)
			return
		}
		StorageQueue.insert(data)
		fmt.Fprintf(responseWriter, `{"success":%t}`, true)
	} else if request.URL.Path[1:] == "setTimeOut" {
		timeOutString := request.Form.Get("timeOut")
		timeOut, err := strconv.Atoi(timeOutString)
		if err != nil {
			fmt.Fprintf(responseWriter, `{"success":%t}`, false)
			return
		}
		StorageQueue.setWaitTime(timeOut)
		fmt.Fprintf(responseWriter, `{"success":%t}`, true)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
