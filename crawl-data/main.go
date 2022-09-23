package main

import (
	"fmt"
	"time"
)

func crawl(worker int, data int) {
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("worker %d is crawling %d \n", worker, data)
}

func main() {
	numberOfData := 1000
	numberOfWorker := 5

	queueChan := make(chan int, numberOfData)
	doneChan := make(chan int)
	for i := 1; i <= numberOfWorker; i++ {
		go func(j int) {
			for v := range queueChan {
				crawl(j, v)
			}
			fmt.Printf("worker %d is done", j)
			doneChan <- j
		}(i)
	}

	for i := 1; i <= numberOfData; i++ {
		queueChan <- i
	}

	close(queueChan)

	for i := 1; i < numberOfWorker; i++ {
		<-doneChan
	}
}
